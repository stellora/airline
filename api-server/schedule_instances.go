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

func syncScheduleInstances(ctx context.Context, queriesTx *db.Queries, schedule db.SchedulesView) error {
	plan, err := planScheduleInstancesSync(ctx, queriesTx, schedule)
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

type syncScheduleAction struct {
	create           *db.CreateFlightInstanceParams
	update           *db.UpdateFlightInstanceParams
	deleteInstanceID int64
}

func planScheduleInstancesSync(ctx context.Context, queriesTx *db.Queries, schedule db.SchedulesView) (plan []syncScheduleAction, err error) {
	scheduleID := sql.NullInt64{Valid: true, Int64: schedule.ID}
	existingInstances, err := queriesTx.ListFlightInstancesForSchedule(ctx, scheduleID)
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
		return existingInstances[i].SourceScheduleInstanceLocaldate.Before(*existingInstances[j].SourceScheduleInstanceLocaldate)
	})

	// Delete any flight instances outside of the schedule start/end dates.
	for _, instance := range existingInstances {
		if instance.SourceScheduleInstanceLocaldate.Before(schedule.StartLocaldate) || instance.SourceScheduleInstanceLocaldate.After(schedule.EndLocaldate) {
			plan = append(plan, syncScheduleAction{deleteInstanceID: instance.ID})
		}
	}

	for curDate := schedule.StartLocaldate; !curDate.After(schedule.EndLocaldate); curDate = curDate.AddDays(1) {
		var curInstance *db.FlightInstancesView
		for _, instance := range existingInstances {
			if curDate.Equal(*instance.SourceScheduleInstanceLocaldate) {
				curInstance = &instance
				break
			}
		}

		hasInstance := curInstance != nil
		shouldHaveInstance := daysOfWeekContains(daysOfWeek, curDate.Weekday())

		if hasInstance && !shouldHaveInstance {
			// Delete the instance.
			plan = append(plan, syncScheduleAction{deleteInstanceID: curInstance.ID})
		} else if !hasInstance && shouldHaveInstance {
			// Create an instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			arrivalDateTime := api.ZonedDateTime{Time: departureDateTime.Add(time.Second * time.Duration(schedule.DurationSec)).In(destinationAirportLoc)}
			plan = append(plan, syncScheduleAction{create: &db.CreateFlightInstanceParams{
				SourceScheduleID:                scheduleID,
				SourceScheduleInstanceLocaldate: &curDate,
				AirlineID:                             schedule.AirlineID,
				Number:                                schedule.Number,
				OriginAirportID:                       schedule.OriginAirportID,
				DestinationAirportID:                  schedule.DestinationAirportID,
				FleetID:                               schedule.FleetID,
				DepartureDatetime:                     &departureDateTime,
				ArrivalDatetime:                       &arrivalDateTime,
				DepartureDatetimeUtc:                  departureDateTime.Time.In(time.UTC),
				ArrivalDatetimeUtc:                    arrivalDateTime.Time.In(time.UTC),
				Published:                             schedule.Published,
			}})
		} else if hasInstance && shouldHaveInstance {
			// Update the existing instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			arrivalDateTime := api.ZonedDateTime{Time: departureDateTime.Add(time.Second * time.Duration(schedule.DurationSec)).In(destinationAirportLoc)}
			plan = append(plan, syncScheduleAction{update: &db.UpdateFlightInstanceParams{
				ID:                   curInstance.ID,
				Number:               sql.NullString{Valid: true, String: schedule.Number},
				OriginAirportID:      sql.NullInt64{Valid: true, Int64: schedule.OriginAirportID},
				DestinationAirportID: sql.NullInt64{Valid: true, Int64: schedule.DestinationAirportID},
				FleetID:              sql.NullInt64{Valid: true, Int64: schedule.FleetID},
				DepartureDatetime:    &departureDateTime,
				ArrivalDatetime:      &arrivalDateTime,
				DepartureDatetimeUtc: sql.NullTime{Valid: true, Time: departureDateTime.Time.In(time.UTC)},
				ArrivalDatetimeUtc:   sql.NullTime{Valid: true, Time: arrivalDateTime.Time.In(time.UTC)},
				Published:            sql.NullBool{Valid: true, Bool: schedule.Published},
			}})
		}
	}

	return plan, nil
}

func (h *Handler) ListFlightInstancesForSchedule(ctx context.Context, request api.ListFlightInstancesForScheduleRequestObject) (api.ListFlightInstancesForScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightInstancesForSchedule(ctx, sql.NullInt64{Valid: true, Int64: int64(request.Id)})
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstancesForSchedule200JSONResponse(mapSlice(fromDBFlightInstance, rows)), nil
}
