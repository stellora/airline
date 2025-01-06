package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightSchedulesByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 CCC-BBB")

	want := api.ListFlightSchedulesByAirport200JSONResponse{
		{
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
			Published:          true},
		{
			Id:                 2,
			Airline:            xxAirline,
			Number:             "2",
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true}}

	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListFlightSchedulesByAirport(ctx, api.ListFlightSchedulesByAirportRequestObject{
			AirportSpec: api.NewAirportSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})

	t.Run("by IATA code", func(t *testing.T) {
		resp, err := handler.ListFlightSchedulesByAirport(ctx, api.ListFlightSchedulesByAirportRequestObject{
			AirportSpec: api.NewAirportSpec(0, "AAA"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})
}
