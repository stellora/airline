package main

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
)

func syncFlightScheduleInstances(ctx context.Context, queriesTx *db.Queries, schedule db.FlightSchedulesView) error {
	plan, err := planFlightScheduleInstancesSync(ctx, queriesTx, schedule)
	if err != nil {
		return err
	}

	for _, action := range plan {
		switch {
		case action.create != nil:
			if _, err := queriesTx.CreateFlightInstance(ctx, *action.create); err != nil {
				return err
			}

		case action.update != nil:
			if _, err := queriesTx.UpdateFlightInstance(ctx, *action.update); err != nil {
				return err
			}

		case action.deleteInstanceID != 0:
			if err := queriesTx.DeleteFlightInstance(ctx, action.deleteInstanceID); err != nil {
				return err
			}
		}
	}

	return nil
}

type syncFlightScheduleAction struct {
	create           *db.CreateFlightInstanceParams
	update           *db.UpdateFlightInstanceParams
	deleteInstanceID int64
}

func planFlightScheduleInstancesSync(ctx context.Context, queriesTx *db.Queries, schedule db.FlightSchedulesView) (plan []syncFlightScheduleAction, err error) {
	scheduleID := sql.NullInt64{Valid: true, Int64: schedule.ID}
	existingInstances, err := queriesTx.ListFlightInstancesForFlightSchedule(ctx, scheduleID)
	if err != nil {
		return nil, err
	}

	originAirportLoc, err := time.LoadLocation(extdata.Airports.AirportByOAID(int(schedule.OriginAirportOadbID.Int64)).Airport.TimezoneID)
	if err != nil {
		return nil, err
	}
	destinationAirportLoc, err := time.LoadLocation(extdata.Airports.AirportByOAID(int(schedule.DestinationAirportOadbID.Int64)).Airport.TimezoneID)
	if err != nil {
		return nil, err
	}

	daysOfWeek, err := parseDaysOfWeek(schedule.DaysOfWeek)
	if err != nil {
		return nil, err
	}

	sort.Slice(existingInstances, func(i, j int) bool {
		return existingInstances[i].SourceFlightScheduleInstanceLocaldate.Before(*existingInstances[j].SourceFlightScheduleInstanceLocaldate)
	})

	// Delete any flight instances outside of the schedule start/end dates.
	for _, instance := range existingInstances {
		if instance.SourceFlightScheduleInstanceLocaldate.Before(schedule.StartLocaldate) || instance.SourceFlightScheduleInstanceLocaldate.After(schedule.EndLocaldate) {
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: instance.ID})
		}
	}

	for curDate := schedule.StartLocaldate; !curDate.After(schedule.EndLocaldate); curDate = curDate.AddDays(1) {
		var curInstance *db.FlightInstancesView
		for _, instance := range existingInstances {
			if curDate.Equal(*instance.SourceFlightScheduleInstanceLocaldate) {
				curInstance = &instance
				break
			}
		}

		hasInstance := curInstance != nil
		shouldHaveInstance := daysOfWeekContains(daysOfWeek, curDate.Weekday())

		if hasInstance && !shouldHaveInstance {
			// Delete the instance.
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: curInstance.ID})
		} else if !hasInstance && shouldHaveInstance {
			// Create an instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			// TODO!(sqs): handle when the arrival is on a different day from the departure
			arrivalDateTime := curDate.TimeOfDay(destinationAirportLoc, schedule.ArrivalLocaltime)
			plan = append(plan, syncFlightScheduleAction{create: &db.CreateFlightInstanceParams{
				SourceFlightScheduleID:                scheduleID,
				SourceFlightScheduleInstanceLocaldate: &curDate,
				AirlineID:                             schedule.AirlineID,
				Number:                                schedule.Number,
				OriginAirportID:                       schedule.OriginAirportID,
				DestinationAirportID:                  schedule.DestinationAirportID,
				AircraftType:                          schedule.AircraftType,
				DepartureDatetime:                     &departureDateTime,
				ArrivalDatetime:                       &arrivalDateTime,
				Published:                             schedule.Published,
			}})
		} else if hasInstance && shouldHaveInstance {
			// Update the existing instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			arrivalDateTime := curDate.TimeOfDay(destinationAirportLoc, schedule.ArrivalLocaltime)
			// TODO!(sqs): handle when the arrival is on a different day from the departure
			plan = append(plan, syncFlightScheduleAction{update: &db.UpdateFlightInstanceParams{
				ID:                   curInstance.ID,
				AirlineID:            sql.NullInt64{Valid: true, Int64: schedule.AirlineID},
				Number:               sql.NullString{Valid: true, String: schedule.Number},
				OriginAirportID:      sql.NullInt64{Valid: true, Int64: schedule.OriginAirportID},
				DestinationAirportID: sql.NullInt64{Valid: true, Int64: schedule.DestinationAirportID},
				AircraftType:         sql.NullString{Valid: true, String: schedule.AircraftType},
				DepartureDatetime:    &departureDateTime,
				ArrivalDatetime:      &arrivalDateTime,
				Published:            sql.NullBool{Valid: true, Bool: schedule.Published},
			}})
		}
	}

	return plan, nil
}

func (h *Handler) ListFlightInstancesForFlightSchedule(ctx context.Context, request api.ListFlightInstancesForFlightScheduleRequestObject) (api.ListFlightInstancesForFlightScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightInstancesForFlightSchedule(ctx, sql.NullInt64{Valid: true, Int64: int64(request.Id)})
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstancesForFlightSchedule200JSONResponse(mapSlice(fromDBFlightInstance, rows)), nil
}
