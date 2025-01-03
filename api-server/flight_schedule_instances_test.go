package main

import (
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stellora/airline/api-server/api"
)

func TestSyncFlightInstancesForFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	flightSchedule1 := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	// Ensure that the flight instances are preserved when updating the schedule.
	setNotesForFlightInstance(t, handler, 1, "a")
	setNotesForFlightInstance(t, handler, 2, "b")
	setNotesForFlightInstance(t, handler, 3, "c")
	checkFlightInstances(t, handler, flightSchedule1.Id, []string{
		"2025-01-01 notes=a",
		"2025-01-02 notes=b",
		"2025-01-03 notes=c",
	})

	_, err := handler.UpdateFlightSchedule(ctx, api.UpdateFlightScheduleRequestObject{
		Id: flightSchedule1.Id,
		Body: &api.UpdateFlightScheduleJSONRequestBody{
			EndDate:    ptrTo(openapi_types.Date{Time: time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)}),
			DaysOfWeek: ptrTo([]int{0, 3, 5, 6}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	checkFlightInstances(t, handler, flightSchedule1.Id, []string{
		"2025-01-01 notes=a",
		"2025-01-03 notes=c",
		"2025-01-04",
		"2025-01-05",
	})
}

func TestListFlightInstancesForFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)
	flightSchedule2 := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	resp, err := handler.ListFlightInstancesForFlightSchedule(ctx, api.ListFlightInstancesForFlightScheduleRequestObject{
		Id: flightSchedule2.Id,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListFlightInstancesForFlightSchedule200JSONResponse{
		{
			Id:           4,
			Source:       flightSchedule2,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:           5,
			Source:       flightSchedule2,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:           6,
			Source:       flightSchedule2,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC)},
		},
	})
}
