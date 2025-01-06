package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/localtime"
)

func TestListFlightSchedulesByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertTestFleet(t, handler)
	insertFleetT(t, handler, "YY", "FF", "")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 AAA-BBB", "YY3 AAA-BBB")

	want := api.ListFlightSchedulesByAirline200JSONResponse{
		api.FlightSchedule{
			Id:                 1,
			Airline:            xxAirline,
			Number:             "1",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true,
		},
		api.FlightSchedule{
			Id:                 2,
			Airline:            xxAirline,
			Number:             "2",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true,
		},
	}

	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListFlightSchedulesByAirline(ctx, api.ListFlightSchedulesByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})

	t.Run("by IATA code", func(t *testing.T) {
		resp, err := handler.ListFlightSchedulesByAirline(ctx, api.ListFlightSchedulesByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(0, "XX"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})
}

func TestListFlightInstancesByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightScheduleT(t, handler,
		fixtureLocalDate1.AddDays(3),
		fixtureLocalDate1.AddDays(4),
		allDaysOfWeek,
	)
	insertFlightInstanceT(t, handler, api.CreateFlightInstanceJSONRequestBody{
		Airline:            api.NewAirlineSpec(0, "XX"),
		Number:             "222",
		OriginAirport:      api.NewAirportSpec(0, "BBB"),
		DestinationAirport: api.NewAirportSpec(0, "AAA"),
		Fleet:              ffFleetSpec,
		DepartureDateTime:  fixtureLocalDate1.TimeOfDay(mustGetTzLocation(bbbAirport.TimezoneID), localtime.NewTimeOfDay(7, 0)),
		ArrivalDateTime:    fixtureLocalDate1.TimeOfDay(mustGetTzLocation(aaaAirport.TimezoneID), localtime.NewTimeOfDay(10, 0)),
		Published:          ptrTo(true),
	})

	want := []string{
		"XX222 BBB-AAA on 2025-01-01",
		"XX1 AAA-BBB on 2025-01-04",
		"XX1 AAA-BBB on 2025-01-05",
	}
	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListFlightInstancesByAirline(ctx, api.ListFlightInstancesByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, flightInstanceDescriptions(resp.(api.ListFlightInstancesByAirline200JSONResponse)), want)
	})

	t.Run("by IATA code", func(t *testing.T) {
		resp, err := handler.ListFlightInstancesByAirline(ctx, api.ListFlightInstancesByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(0, "XX"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, flightInstanceDescriptions(resp.(api.ListFlightInstancesByAirline200JSONResponse)), want)
	})
}
