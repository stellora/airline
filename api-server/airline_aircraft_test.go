package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListAircraftByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertAircraftT(t, handler, "XX", "N1", "N2")
	insertAircraftT(t, handler, "YY", "N3")

	want := api.ListAircraftByAirline200JSONResponse{
		{Id: 1, Registration: "N1", AircraftType: "777", Airline: api.Airline{Id: 1, IataCode: "XX"}},
		{Id: 2, Registration: "N2", AircraftType: "777", Airline: api.Airline{Id: 1, IataCode: "XX"}},
	}

	t.Run("by id", func(t *testing.T) {
		resp, err := handler.ListAircraftByAirline(ctx, api.ListAircraftByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(1, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})

	t.Run("by IATA code", func(t *testing.T) {
		resp, err := handler.ListAircraftByAirline(ctx, api.ListAircraftByAirlineRequestObject{
			AirlineSpec: api.NewAirlineSpec(0, "XX"),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, want)
	})
}
