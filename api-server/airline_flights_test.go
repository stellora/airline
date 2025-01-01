package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightsByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertFlightsT(t, handler, "XX1 AAA-BBB", "XX2 AAA-BBB", "YY3 AAA-BBB")

	want := api.ListFlightsByAirline200JSONResponse{
		{Id: 1, Number: "XX1", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}, Published: true},
		{Id: 2, Number: "XX2", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}, Published: true},
	}

	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListFlightsByAirline(ctx, api.ListFlightsByAirlineRequestObject{
			AirlineSpec: newAirlineSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})

	t.Run("by airline code", func(t *testing.T) {
		resp, err := handler.ListFlightsByAirline(ctx, api.ListFlightsByAirlineRequestObject{
			AirlineSpec: newAirlineSpec(0, "XX"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})
}
