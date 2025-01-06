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

var fixtureManualFlightInstance = api.CreateFlightInstanceJSONRequestBody{
	Airline:            api.NewAirlineSpec(0, "XX"),
	Number:             "222",
	OriginAirport:      api.NewAirportSpec(0, "AAA"),
	DestinationAirport: api.NewAirportSpec(0, "BBB"),
	AircraftType:       fixtureB77W.IcaoCode,
	DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
	ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
	Published:          ptrTo(true),
}

func TestGetFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	flightSchedule := insertFlightScheduleT(t, handler,
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
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
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(1).String()),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
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
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
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
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.String()),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            flightSchedule.Published,
		},
		{
			Id:                   2,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(1).String()),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            flightSchedule.Published,
		},
		{
			Id:                   3,
			ScheduleID:           &flightSchedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(2).String()),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    fixtureLocalDate1.AddDays(2).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(2).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            flightSchedule.Published,
		},
	})
}

func TestCreateFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	resp, err := handler.CreateFlightInstance(ctx, api.CreateFlightInstanceRequestObject{
		Body: &api.CreateFlightInstanceJSONRequestBody{
			Airline:            api.NewAirlineSpec(0, "XX"),
			Number:             "222",
			OriginAirport:      api.NewAirportSpec(0, "AAA"),
			DestinationAirport: api.NewAirportSpec(0, "BBB"),
			AircraftType:       fixtureB77W.IcaoCode,
			DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
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
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
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
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.String()),
			Airline:              flightSchedule.Airline,
			Number:               flightSchedule.Number,
			OriginAirport:        flightSchedule.OriginAirport,
			DestinationAirport:   flightSchedule.DestinationAirport,
			AircraftType:         flightSchedule.AircraftType,
			DepartureDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            flightSchedule.Published,
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
		DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
		ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
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

func flightInstanceDescription(instance api.FlightInstance) string {
	parts := []string{fmt.Sprintf("%s%s %s-%s on %s", instance.Airline.IataCode, instance.Number, instance.OriginAirport.IataCode, instance.DestinationAirport.IataCode, instance.DepartureDateTime.Format(time.DateOnly))}
	if instance.Aircraft != nil {
		parts = append(parts, fmt.Sprintf("aircraft=%s", instance.Aircraft.AircraftType))
	}
	if instance.Notes != "" {
		parts = append(parts, fmt.Sprintf("notes=%s", instance.Notes))
	}
	return strings.Join(parts, " ")
}

func flightInstanceDescriptions(flights []api.FlightInstance) []string {
	return mapSlice(flightInstanceDescription, flights)
}

func checkFlightInstances(t *testing.T, handler *Handler, flightScheduleID int, want []string) {
	t.Helper()

	resp, err := handler.ListFlightInstancesForFlightSchedule(context.Background(), api.ListFlightInstancesForFlightScheduleRequestObject{Id: flightScheduleID})
	if err != nil {
		t.Fatal(err)
	}

	instances := resp.(api.ListFlightInstancesForFlightSchedule200JSONResponse)
	assertEqual(t, flightInstanceDescriptions(instances), want)
}
