package main

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
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

	daysOfWeek, err := parseDaysOfWeek(schedule.DaysOfWeek)
	if err != nil {
		return nil, err
	}

	sort.Slice(existingInstances, func(i, j int) bool {
		return existingInstances[i].SourceFlightScheduleInstanceDate.Time.Before(existingInstances[j].SourceFlightScheduleInstanceDate.Time)
	})

	// Delete any flight instances outside of the schedule start/end dates.
	for _, instance := range existingInstances {
		if instance.SourceFlightScheduleInstanceDate.Time.Before(schedule.StartDate) || instance.SourceFlightScheduleInstanceDate.Time.After(schedule.EndDate) {
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: instance.ID})
		}
	}

	for curDate := schedule.StartDate; !curDate.After(schedule.EndDate); curDate = curDate.AddDate(0, 0, 1) {
		var curInstance *db.FlightInstancesView
		for _, instance := range existingInstances {
			if curDate.Equal(instance.SourceFlightScheduleInstanceDate.Time) {
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
			departureDateTime, err := scheduleToInstanceDateTime(curDate, schedule.DepartureTime)
			if err != nil {
				return nil, err
			}
			// TODO!(sqs): handle when the arrival is on a different day from the departure
			arrivalDateTime, err := scheduleToInstanceDateTime(curDate, schedule.ArrivalTime)
			if err != nil {
				return nil, err
			}
			plan = append(plan, syncFlightScheduleAction{create: &db.CreateFlightInstanceParams{
				SourceFlightScheduleID:           scheduleID,
				SourceFlightScheduleInstanceDate: sql.NullTime{Valid: true, Time: curDate},
				AirlineID:                        schedule.AirlineID,
				Number:                           schedule.Number,
				OriginAirportID:                  schedule.OriginAirportID,
				DestinationAirportID:             schedule.DestinationAirportID,
				AircraftType:                     schedule.AircraftType,
				DepartureDatetime:                departureDateTime,
				ArrivalDatetime:                  arrivalDateTime,
				Published:                        schedule.Published,
			}})
		} else if hasInstance && shouldHaveInstance {
			// Update the existing instance.
			departureDateTime, err := scheduleToInstanceDateTime(curDate, schedule.DepartureTime)
			if err != nil {
				return nil, err
			}
			// TODO!(sqs): handle when the arrival is on a different day from the departure
			arrivalDateTime, err := scheduleToInstanceDateTime(curDate, schedule.ArrivalTime)
			if err != nil {
				return nil, err
			}
			plan = append(plan, syncFlightScheduleAction{update: &db.UpdateFlightInstanceParams{
				ID:                   curInstance.ID,
				AirlineID:            sql.NullInt64{Valid: true, Int64: schedule.AirlineID},
				Number:               sql.NullString{Valid: true, String: schedule.Number},
				OriginAirportID:      sql.NullInt64{Valid: true, Int64: schedule.OriginAirportID},
				DestinationAirportID: sql.NullInt64{Valid: true, Int64: schedule.DestinationAirportID},
				AircraftType:         sql.NullString{Valid: true, String: schedule.AircraftType},
				DepartureDatetime:    sql.NullTime{Valid: true, Time: departureDateTime},
				ArrivalDatetime:      sql.NullTime{Valid: true, Time: arrivalDateTime},
				Published:            sql.NullBool{Valid: true, Bool: schedule.Published},
			}})
		}
	}

	return plan, nil
}

func scheduleToInstanceDateTime(instanceDate time.Time, timeOfDay string, loc *time.Location) (time.Time, error) {
	const HourMinuteOnly = "15:04"
	tm, err := time.Parse(HourMinuteOnly, timeOfDay)
	if err != nil {
		return time.Time{}, err
	}

	year, month, day := instanceDate.Date()
	// TODO!(sqs): use timezone of airport
	return time.Date(year, month, day, tm.Hour(), tm.Minute(), 0, 0, time.UTC), nil
}

func (h *Handler) ListFlightInstancesForFlightSchedule(ctx context.Context, request api.ListFlightInstancesForFlightScheduleRequestObject) (api.ListFlightInstancesForFlightScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightInstancesForFlightSchedule(ctx, sql.NullInt64{Valid: true, Int64: int64(request.Id)})
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstancesForFlightSchedule200JSONResponse(mapSlice(fromDBFlightInstance, rows)), nil
}
