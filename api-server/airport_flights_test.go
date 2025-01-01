package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightsByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightsT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 CCC-BBB")

	want := api.ListFlightsByAirport200JSONResponse{
		{Id: 1, Number: "XX1", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}, Published: true},
		{Id: 2, Number: "XX2", OriginAirport: api.Airport{Id: 2, IataCode: "BBB"}, DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"}, Published: true},
	}

	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListFlightsByAirport(ctx, api.ListFlightsByAirportRequestObject{
			AirportSpec: newAirportSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})

	t.Run("by airport code", func(t *testing.T) {
		resp, err := handler.ListFlightsByAirport(ctx, api.ListFlightsByAirportRequestObject{
			AirportSpec: newAirportSpec(0, "AAA"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})
}
