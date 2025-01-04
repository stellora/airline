package main

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/localtime"
)

func TestGetFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	flightSchedule := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
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
			ScheduleInstanceDate: ptrTo("2025-01-02"),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    time.Date(2025, 1, 2, 7, 0, 0, 0, time.UTC),
			ArrivalDateTime:      time.Date(2025, 1, 2, 9, 0, 0, 0, time.UTC),
			Published:            flightSchedule.Published,
		})
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.GetFlightInstance404Response{})
	})
}

func TestListFlightInstances(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	flightSchedule := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
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
			ScheduleInstanceDate: ptrTo("2025-01-01"),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    time.Date(2025, 1, 1, 7, 0, 0, 0, time.UTC),
			ArrivalDateTime:      time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC),
			Published:            flightSchedule.Published,
		},
		{
			Id:                   2,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: ptrTo("2025-01-02"),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    time.Date(2025, 1, 2, 7, 0, 0, 0, time.UTC),
			ArrivalDateTime:      time.Date(2025, 1, 2, 9, 0, 0, 0, time.UTC),
			Published:            flightSchedule.Published,
		},
		{
			Id:                   3,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: ptrTo("2025-01-03"),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    time.Date(2025, 1, 3, 7, 0, 0, 0, time.UTC),
			ArrivalDateTime:      time.Date(2025, 1, 3, 9, 0, 0, 0, time.UTC),
			Published:            flightSchedule.Published,
		},
	})
}

func TestCreateFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	aaaTz, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Fatal(err)
	}
	bbbTz, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := handler.CreateFlightInstance(ctx, api.CreateFlightInstanceRequestObject{
		Body: &api.CreateFlightInstanceJSONRequestBody{
			Airline:            api.NewAirlineSpec(0, "XX"),
			Number:             "222",
			OriginAirport:      api.NewAirportSpec(0, "AAA"),
			DestinationAirport: api.NewAirportSpec(0, "BBB"),
			AircraftType:       fixtureB77W.IcaoCode,
			DepartureDateTime:  fixtureLocalDate1.TimeOfDay(aaaTz, localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(bbbTz, localtime.NewTimeOfDay(10, 0)),
			Published:          ptrTo(true),
		}})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp.(api.CreateFlightInstance201JSONResponse); !ok {
		t.Fatalf("got %T, want non-error response", resp)
	}
}

func TestUpdateFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	flightSchedule := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
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
			ScheduleInstanceDate: ptrTo("2025-01-01"),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    time.Date(2025, 1, 1, 7, 0, 0, 0, time.UTC),
			ArrivalDateTime:      time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC),
			Published:            flightSchedule.Published,
			Notes:                "abc",
		})
	}
}

func TestDeleteFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	aaaTz, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Fatal(err)
	}
	bbbTz, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}

	manualFlightInstance := insertFlightInstanceT(t, handler, api.CreateFlightInstanceJSONRequestBody{
		Airline:            api.NewAirlineSpec(0, "XX"),
		Number:             "222",
		OriginAirport:      api.NewAirportSpec(0, "AAA"),
		DestinationAirport: api.NewAirportSpec(0, "BBB"),
		AircraftType:       fixtureB77W.IcaoCode,
		DepartureDateTime:  fixtureLocalDate1.TimeOfDay(aaaTz, localtime.NewTimeOfDay(7, 0)),
		ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(bbbTz, localtime.NewTimeOfDay(10, 0)),
		Published:          ptrTo(true),
	})
	flightSchedule := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)

	checkFlightInstanceExistence := func(t *testing.T, id int, wantExist bool) {
		t.Helper()
		resp, err := handler.GetFlightInstance(ctx, api.GetFlightInstanceRequestObject{Id: id})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("AAA %#v", resp)
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
		assertEqual(t, resp, api.DeleteFlightInstance400Response{})
		checkFlightInstanceExistence(t, instances[0].Id, true)
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
		parts := []string{*instance.ScheduleInstanceDate}
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
