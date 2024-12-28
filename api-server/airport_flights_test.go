package main

import (
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListFlightsByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
		{Id: 2, Number: "ST2", OriginAirport: api.Airport{Id: 2}, DestinationAirport: api.Airport{Id: 1}},
		{Id: 3, Number: "ST3", OriginAirport: api.Airport{Id: 3}, DestinationAirport: api.Airport{Id: 2}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
		{Id: 3, IataCode: "CCC"},
	}

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
