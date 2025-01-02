package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightSchedulesByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 AAA-BBB", "YY3 AAA-BBB")

	want := api.ListFlightSchedulesByAirline200JSONResponse{
		{Id: 1, Airline: api.Airline{Id: 1, IataCode: "XX"}, Number: "1", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}, Published: true},
		{Id: 2, Airline: api.Airline{Id: 1, IataCode: "XX"}, Number: "2", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}, Published: true},
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
