package main

import (
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListRoutes(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertFlightsT(t, handler, "ST1 AAA-BBB", "ST2 BBB-AAA", "ST3 AAA-CCC", "ST4 AAA-CCC")

	resp, err := handler.ListRoutes(ctx, api.ListRoutesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListRoutes200JSONResponse{
		{
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 3, IataCode: "CCC"},
			FlightsCount:       2,
		},
		{
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
			FlightsCount:       1,
		},
		{
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
			FlightsCount:       1,
		},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}
