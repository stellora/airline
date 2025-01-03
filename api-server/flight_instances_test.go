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
			Id:                   2,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: &openapi_types.Date{Time: time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
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
			Id:                   1,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: &openapi_types.Date{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:                   2,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: &openapi_types.Date{Time: time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			Id:                   3,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: &openapi_types.Date{Time: time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC)},
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
			t.Fatalf("got %T, want non-error response", resp)
		}
		if got := resp.(api.UpdateFlightInstance200JSONResponse); got.Notes != notes {
			t.Errorf("got notes %q, want %q", got.Notes, notes)
		}
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, resp, api.GetFlightInstance200JSONResponse{
			Id:                   1,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: &openapi_types.Date{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
			Notes:                "abc",
		})
	}
}

func TestDeleteFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	manualFlightInstance := insertFlightInstanceT(t, handler, api.CreateFlightInstanceJSONRequestBody{
		Airline:            api.NewAirlineSpec(0, "XX"),
		Number:             "222",
		OriginAirport:      api.NewAirportSpec(0, "AAA"),
		DestinationAirport: api.NewAirportSpec(0, "BBB"),
		AircraftType:       fixtureB77W.IcaoCode,
		DepartureDateTime:  fixtureDate1,
		ArrivalDateTime:    openapi_types.Date{Time: fixtureDate1.Time.Add(3 * time.Hour)},
		Published:          ptrTo(true),
	})
	flightSchedule := insertFlightScheduleT(t, handler,
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		allDaysOfWeek,
	)

	checkFlightInstanceExistence := func(t *testing.T, id int, wantExist bool) {
		t.Helper()
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{Id: id})
		if err != nil {
			t.Fatal(err)
		}
		if _, isNotExist := resp.(api.GetFlightInstance404Response); !isNotExist != wantExist {
			t.Fatalf("flight instance %d: got exists %v, want %v", id, !isNotExist, wantExist)
		}
	}

	t.Run("source = FlightSchedule", func(t *testing.T) {
		instancesResp, err := handler.ListFlightInstancesForFlightSchedule(ctx, api.ListFlightInstancesForFlightScheduleRequestObject{Id: flightSchedule.Id})
		if err != nil {
			t.Fatal(err)
		}
		instances := instancesResp.(api.ListFlightInstancesForFlightSchedule200JSONResponse)

		checkFlightInstanceExistence(t, instances[0].Id, true)
		resp, err := handler.DeleteFlightInstance(ctx, api.DeleteFlightInstanceRequestObject{Id: instances[0].Id})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.DeleteFlightInstance204Response{})
		checkFlightInstanceExistence(t, instances[0].Id, false)
	})

	t.Run("source = manual", func(t *testing.T) {
		checkFlightInstanceExistence(t, manualFlightInstance.Id, true)
		resp, err := handler.DeleteFlightInstance(ctx, api.DeleteFlightInstanceRequestObject{Id: manualFlightInstance.Id})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.DeleteFlightInstance204Response{})
		checkFlightInstanceExistence(t, manualFlightInstance.Id, false)
	})
}

func checkFlightInstances(t *testing.T, handler *Handler, flightScheduleID int, want []string) {
	t.Helper()

	resp, err := handler.ListFlightInstancesForFlightSchedule(context.Background(), api.ListFlightInstancesForFlightScheduleRequestObject{Id: flightScheduleID})
	if err != nil {
		t.Fatal(err)
	}

	toDescription := func(instance api.FlightInstance) string {
		parts := []string{instance.ScheduleInstanceDate.Time.Format("2006-01-02")}
		if instance.Aircraft != nil {
			parts = append(parts, fmt.Sprintf("aircraft=%s", instance.Aircraft.AircraftType))
		}
		if instance.Notes != "" {
			parts = append(parts, fmt.Sprintf("notes=%s", instance.Notes))
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
