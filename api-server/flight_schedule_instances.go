package main

import (
	"context"
	"sort"

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
	existingInstances, err := queriesTx.ListFlightInstancesForFlightSchedule(ctx, schedule.ID)
	if err != nil {
		return nil, err
	}

	if !schedule.Published {
		// Unpublished schedules have no flight instances.
		for _, instance := range existingInstances {
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: instance.FlightInstance.ID})
		}
		return plan, nil
	}

	daysOfWeek, err := parseDaysOfWeek(schedule.DaysOfWeek)
	if err != nil {
		return nil, err
	}

	sort.Slice(existingInstances, func(i, j int) bool {
		return existingInstances[i].FlightInstance.InstanceDate.Before(existingInstances[j].FlightInstance.InstanceDate)
	})

	// Delete any flight instances outside of the schedule start/end dates.
	for _, instance := range existingInstances {
		if instance.FlightInstance.InstanceDate.Before(schedule.StartDate) || instance.FlightInstance.InstanceDate.After(schedule.EndDate) {
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: instance.FlightInstance.ID})
		}
	}

	for curDate := schedule.StartDate; !curDate.After(schedule.EndDate); curDate = curDate.AddDate(0, 0, 1) {
		var curInstance *db.ListFlightInstancesForFlightScheduleRow
		for _, instance := range existingInstances {
			if curDate.Equal(instance.FlightInstance.InstanceDate) {
				curInstance = &instance
				break
			}
		}

		hasInstance := curInstance != nil
		shouldHaveInstance := daysOfWeekContains(daysOfWeek, curDate.Weekday())

		if hasInstance && !shouldHaveInstance {
			// Delete the instance.
			plan = append(plan, syncFlightScheduleAction{deleteInstanceID: curInstance.FlightInstance.ID})
		} else if !hasInstance && shouldHaveInstance {
			// Create an instance.
			plan = append(plan, syncFlightScheduleAction{create: &db.CreateFlightInstanceParams{
				SourceFlightScheduleID: schedule.ID,
				InstanceDate:           curDate,
			}})
		}
	}

	return plan, nil
}

func (h *Handler) ListFlightInstancesForFlightSchedule(ctx context.Context, request api.ListFlightInstancesForFlightScheduleRequestObject) (api.ListFlightInstancesForFlightScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightInstancesForFlightSchedule(ctx, int64(request.Id))
	if err != nil {
		return nil, err
	}

	rows2 := make([]db.ListFlightInstancesRow, len(rows))
	for i, r := range rows {
		rows2[i] = db.ListFlightInstancesRow(r)
	}
	return api.ListFlightInstancesForFlightSchedule200JSONResponse(fromDBFlightInstances(rows2)), nil
}
