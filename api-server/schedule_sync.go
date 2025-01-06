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

func syncScheduleFlightInstances(ctx context.Context, queriesTx *db.Queries, schedule db.SchedulesView) error {
	plan, err := planScheduleFlightInstancesSync(ctx, queriesTx, schedule)
	if err != nil {
		return err
	}

	for _, action := range plan {
		switch {
		case action.create != nil:
			if _, err := queriesTx.CreateFlight(ctx, *action.create); err != nil {
				return err
			}

		case action.update != nil:
			if _, err := queriesTx.UpdateFlight(ctx, *action.update); err != nil {
				return err
			}

		case action.deleteFlightID != 0:
			if err := queriesTx.DeleteFlight(ctx, action.deleteFlightID); err != nil {
				return err
			}
		}
	}

	return nil
}

type syncScheduleAction struct {
	create         *db.CreateFlightParams
	update         *db.UpdateFlightParams
	deleteFlightID int64
}

func planScheduleFlightInstancesSync(ctx context.Context, queriesTx *db.Queries, schedule db.SchedulesView) (plan []syncScheduleAction, err error) {
	scheduleID := sql.NullInt64{Valid: true, Int64: schedule.ID}
	existingFlights, err := queriesTx.ListFlightsForSchedule(ctx, scheduleID)
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

	sort.Slice(existingFlights, func(i, j int) bool {
		return existingFlights[i].SourceScheduleInstanceLocaldate.Before(*existingFlights[j].SourceScheduleInstanceLocaldate)
	})

	// Delete any flights outside of the schedule start/end dates.
	for _, instance := range existingFlights {
		if instance.SourceScheduleInstanceLocaldate.Before(schedule.StartLocaldate) || instance.SourceScheduleInstanceLocaldate.After(schedule.EndLocaldate) {
			plan = append(plan, syncScheduleAction{deleteFlightID: instance.ID})
		}
	}

	for curDate := schedule.StartLocaldate; !curDate.After(schedule.EndLocaldate); curDate = curDate.AddDays(1) {
		var curInstance *db.FlightsView
		for _, instance := range existingFlights {
			if curDate.Equal(*instance.SourceScheduleInstanceLocaldate) {
				curInstance = &instance
				break
			}
		}

		hasInstance := curInstance != nil
		shouldHaveInstance := daysOfWeekContains(daysOfWeek, curDate.Weekday())

		if hasInstance && !shouldHaveInstance {
			// Delete the instance.
			plan = append(plan, syncScheduleAction{deleteFlightID: curInstance.ID})
		} else if !hasInstance && shouldHaveInstance {
			// Create an instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			arrivalDateTime := api.ZonedDateTime{Time: departureDateTime.Add(time.Second * time.Duration(schedule.DurationSec)).In(destinationAirportLoc)}
			plan = append(plan, syncScheduleAction{create: &db.CreateFlightParams{
				SourceScheduleID:                scheduleID,
				SourceScheduleInstanceLocaldate: &curDate,
				AirlineID:                       schedule.AirlineID,
				Number:                          schedule.Number,
				OriginAirportID:                 schedule.OriginAirportID,
				DestinationAirportID:            schedule.DestinationAirportID,
				FleetID:                         schedule.FleetID,
				DepartureDatetime:               &departureDateTime,
				ArrivalDatetime:                 &arrivalDateTime,
				DepartureDatetimeUtc:            departureDateTime.Time.In(time.UTC),
				ArrivalDatetimeUtc:              arrivalDateTime.Time.In(time.UTC),
				Published:                       schedule.Published,
			}})
		} else if hasInstance && shouldHaveInstance {
			// Update the existing instance.
			departureDateTime := curDate.TimeOfDay(originAirportLoc, schedule.DepartureLocaltime)
			arrivalDateTime := api.ZonedDateTime{Time: departureDateTime.Add(time.Second * time.Duration(schedule.DurationSec)).In(destinationAirportLoc)}
			plan = append(plan, syncScheduleAction{update: &db.UpdateFlightParams{
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

func (h *Handler) ListFlightsForSchedule(ctx context.Context, request api.ListFlightsForScheduleRequestObject) (api.ListFlightsForScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightsForSchedule(ctx, sql.NullInt64{Valid: true, Int64: int64(request.Id)})
	if err != nil {
		return nil, err
	}
	return api.ListFlightsForSchedule200JSONResponse(mapSlice(fromDBFlight, rows)), nil
}
