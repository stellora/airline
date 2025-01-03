package main

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stellora/airline/api-server/api"
)

func TestGetFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	flightSchedule := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{
			Id: 2,
		})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, resp, api.GetFlightInstance200JSONResponse{
			Id:           2,
			Source:       flightSchedule,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
		})
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetFlightInstance404Response{})
	})
}

func TestListFlightInstances(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	flightSchedule := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	resp, err := handler.ListFlightInstances(ctx, api.ListFlightInstancesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListFlightInstances200JSONResponse{
		{
			Id:           1,
			Source:       flightSchedule,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:           2,
			Source:       flightSchedule,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:           3,
			Source:       flightSchedule,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC)},
		},
	})
}

func TestUpdateFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	flightSchedule := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	{
		// Update the flight
		notes := "abc"
		resp, err := handler.UpdateFlightInstance(ctx, api.UpdateFlightInstanceRequestObject{
			Id: 1,
			Body: &api.UpdateFlightInstanceJSONRequestBody{
				Notes: ptrTo(notes),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.UpdateFlightInstance200JSONResponse); !ok {
			t.Errorf("got %#v", resp)
		}
		if got := resp.(api.UpdateFlightInstance200JSONResponse); *got.Notes != notes {
			t.Errorf("got notes %q, want %q", *got.Notes, notes)
		}
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, resp, api.GetFlightInstance200JSONResponse{
			Id:           1,
			Source:       flightSchedule,
			InstanceDate: openapi_types.Date{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
			Notes:        ptrTo("abc"),
		})
	}
}

func checkFlightInstances(t *testing.T, handler *Handler, flightScheduleID int, want []string) {
	t.Helper()

	resp, err := handler.ListFlightInstancesForFlightSchedule(context.Background(), api.ListFlightInstancesForFlightScheduleRequestObject{Id: flightScheduleID})
	if err != nil {
		t.Fatal(err)
	}

	toDescription := func(instance api.FlightInstance) string {
		parts := []string{instance.InstanceDate.Time.Format("2006-01-02")}
		if instance.Aircraft != nil {
			parts = append(parts, fmt.Sprintf("aircraft=%s", instance.Aircraft.AircraftType))
		}
		if instance.Notes != nil {
			parts = append(parts, fmt.Sprintf("notes=%s", *instance.Notes))
		}
		return strings.Join(parts, " ")
	}

	instances := resp.(api.ListFlightInstancesForFlightSchedule200JSONResponse)
	descriptions := make([]string, len(instances))
	for i, instance := range instances {
		descriptions[i] = toDescription(instance)
	}

	assertEqual(t, descriptions, want)
}
