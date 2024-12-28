package main

import (
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightsByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler.queries, "AAA", "BBB", "CCC")
	insertFlightsT(t, handler.queries, "ST1 AAA-BBB", "ST2 BBB-AAA", "ST3 CCC-BBB")

	resp, err := handler.ListFlightsByAirport(ctx, api.ListFlightsByAirportRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlightsByAirport200JSONResponse{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1, IataCode: "AAA"}, DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"}},
		{Id: 2, Number: "ST2", OriginAirport: api.Airport{Id: 2, IataCode: "BBB"}, DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"}},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}
