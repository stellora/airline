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

var fixtureManualFlight = api.CreateFlightJSONRequestBody{
	Airline:            api.NewAirlineSpec(0, "XX"),
	Number:             "222",
	OriginAirport:      api.NewAirportSpec(0, "AAA"),
	DestinationAirport: api.NewAirportSpec(0, "BBB"),
	Fleet:              ffFleetSpec,
	Aircraft:           nil,
	DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
	ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
	Published:          ptrTo(true),
}

func TestGetFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	schedule := insertScheduleT(t, handler,
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
		allDaysOfWeek,
	)

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 2,
		})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, resp, api.GetFlight200JSONResponse{
			Id:                   2,
			ScheduleID:           &schedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(1).String()),
			Airline:              schedule.Airline,
			Number:               schedule.Number,
			OriginAirport:        schedule.OriginAirport,
			DestinationAirport:   schedule.DestinationAirport,
			Fleet:                schedule.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            schedule.Published,
		})
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.GetFlight404Response{})
	})
}

func TestListFlights(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	schedule := insertScheduleT(t, handler,
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
		allDaysOfWeek,
	)

	resp, err := handler.ListFlights(ctx, api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListFlights200JSONResponse{
		{
			Id:                   1,
			ScheduleID:           &schedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.String()),
			Airline:              schedule.Airline,
			Number:               schedule.Number,
			OriginAirport:        schedule.OriginAirport,
			DestinationAirport:   schedule.DestinationAirport,
			Fleet:                schedule.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            schedule.Published,
		},
		{
			Id:                   2,
			ScheduleID:           &schedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(1).String()),
			Airline:              schedule.Airline,
			Number:               schedule.Number,
			OriginAirport:        schedule.OriginAirport,
			DestinationAirport:   schedule.DestinationAirport,
			Fleet:                schedule.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(1).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            schedule.Published,
		},
		{
			Id:                   3,
			ScheduleID:           &schedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.AddDays(2).String()),
			Airline:              schedule.Airline,
			Number:               schedule.Number,
			OriginAirport:        schedule.OriginAirport,
			DestinationAirport:   schedule.DestinationAirport,
			Fleet:                schedule.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    fixtureLocalDate1.AddDays(2).TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.AddDays(2).TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            schedule.Published,
		},
	})
}

func TestCreateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertTestFleet(t, handler)

	resp, err := handler.CreateFlight(ctx, api.CreateFlightRequestObject{
		Body: &api.CreateFlightJSONRequestBody{
			Airline:            api.NewAirlineSpec(0, "XX"),
			Number:             "222",
			OriginAirport:      api.NewAirportSpec(0, "AAA"),
			DestinationAirport: api.NewAirportSpec(0, "BBB"),
			Fleet:              ffFleetSpec,
			Aircraft:           nil,
			DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
			Published:          ptrTo(true),
		}})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp.(api.CreateFlight201JSONResponse); !ok {
		t.Fatalf("got %T, want non-error response", resp)
	}
}

func TestUpdateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	schedule := insertScheduleT(t, handler,
		fixtureLocalDate1,
		fixtureLocalDate1.AddDays(2),
		allDaysOfWeek,
	)

	{
		// Update the flight
		notes := "abc"
		resp, err := handler.UpdateFlight(ctx, api.UpdateFlightRequestObject{
			Id: 1,
			Body: &api.UpdateFlightJSONRequestBody{
				Notes: ptrTo(notes),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.UpdateFlight200JSONResponse); !ok {
			t.Fatalf("got %T, want non-error response", resp)
		}
		if got := resp.(api.UpdateFlight200JSONResponse); got.Notes != notes {
			t.Errorf("got notes %q, want %q", got.Notes, notes)
		}
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, resp, api.GetFlight200JSONResponse{
			Id:                   1,
			ScheduleID:           &schedule.Id,
			ScheduleInstanceDate: ptrTo(fixtureLocalDate1.String()),
			Airline:              schedule.Airline,
			Number:               schedule.Number,
			OriginAirport:        schedule.OriginAirport,
			DestinationAirport:   schedule.DestinationAirport,
			Fleet:                schedule.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
			ArrivalDateTime:      fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(12, 0)),
			Published:            schedule.Published,
			Notes:                "abc",
		})
	}
}

func TestDeleteFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	manualFlight := insertFlightT(t, handler, api.CreateFlightJSONRequestBody{
		Airline:            api.NewAirlineSpec(0, "XX"),
		Number:             "222",
		OriginAirport:      api.NewAirportSpec(0, "AAA"),
		DestinationAirport: api.NewAirportSpec(0, "BBB"),
		Fleet:              ffFleetSpec,
		DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
		ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
		Published:          ptrTo(true),
	})
	schedule := insertScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)

	checkFlightExistence := func(t *testing.T, id int, wantExist bool) {
		t.Helper()
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{Id: id})
		if err != nil {
			t.Fatal(err)
		}
		if _, isNotExist := resp.(api.GetFlight404Response); !isNotExist != wantExist {
			t.Fatalf("flight %d: got exists %v, want %v", id, !isNotExist, wantExist)
		}
	}

	t.Run("source = Schedule", func(t *testing.T) {
		flightsResp, err := handler.ListFlightsForSchedule(ctx, api.ListFlightsForScheduleRequestObject{Id: schedule.Id})
		if err != nil {
			t.Fatal(err)
		}
		flights := flightsResp.(api.ListFlightsForSchedule200JSONResponse)

		checkFlightExistence(t, flights[0].Id, true)
		resp, err := handler.DeleteFlight(ctx, api.DeleteFlightRequestObject{Id: flights[0].Id})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.DeleteFlight400Response{})
		checkFlightExistence(t, flights[0].Id, true)
	})

	t.Run("source = manual", func(t *testing.T) {
		checkFlightExistence(t, manualFlight.Id, true)
		resp, err := handler.DeleteFlight(ctx, api.DeleteFlightRequestObject{Id: manualFlight.Id})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.DeleteFlight204Response{})
		checkFlightExistence(t, manualFlight.Id, false)
	})
}

func flightDescription(f api.Flight) string {
	parts := []string{fmt.Sprintf("%s%s %s-%s on %s", f.Airline.IataCode, f.Number, f.OriginAirport.IataCode, f.DestinationAirport.IataCode, f.DepartureDateTime.Format(time.DateOnly))}
	if f.Aircraft != nil {
		parts = append(parts, fmt.Sprintf("aircraft=%s", f.Aircraft.AircraftType))
	}
	if f.Notes != "" {
		parts = append(parts, fmt.Sprintf("notes=%s", f.Notes))
	}
	return strings.Join(parts, " ")
}

func flightDescriptions(flights []api.Flight) []string {
	return mapSlice(flightDescription, flights)
}

func checkFlights(t *testing.T, handler *Handler, scheduleID int, want []string) {
	t.Helper()

	resp, err := handler.ListFlightsForSchedule(context.Background(), api.ListFlightsForScheduleRequestObject{Id: scheduleID})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, flightDescriptions(resp.(api.ListFlightsForSchedule200JSONResponse)), want)
}
