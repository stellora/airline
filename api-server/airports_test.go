package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetAirport200JSONResponse{
			Id:       1,
			IataCode: "AAA",
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetAirport404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})
}

func TestListAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []*api.Airport{
		{IataCode: "AAA"},
		{IataCode: "BBB"},
	}

	resp, err := handler.ListAirports(ctx, api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAirports200JSONResponse{
		api.Airport{IataCode: "AAA"},
		api.Airport{IataCode: "BBB"},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateAirport(t *testing.T) {
	ctx, handler := handlerTest(t)

	resp, err := handler.CreateAirport(ctx, api.CreateAirportRequestObject{
		Body: &api.CreateAirportJSONRequestBody{
			IataCode: "New Airport",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateAirport201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportIataCodes(t, handler, []string{"New Airport"})
}

func TestDeleteAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	resp, err := handler.DeleteAirport(ctx, api.DeleteAirportRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirport204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportIataCodes(t, handler, []string{"BBB"})
}

func TestDeleteAllAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []*api.Airport{
		{IataCode: "AAA"},
		{IataCode: "BBB"},
	}

	resp, err := handler.DeleteAllAirports(ctx, api.DeleteAllAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllAirports204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportIataCodes(t, handler, []string{})
}

func TestListFlightsByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
		{Id: 2, Number: "ST2", OriginAirport: api.Airport{Id: 2}, DestinationAirport: api.Airport{Id: 1}},
		{Id: 3, Number: "ST3", OriginAirport: api.Airport{Id: 3}, DestinationAirport: api.Airport{Id: 4}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	resp, err := handler.ListFlightsByAirport(ctx, api.ListFlightsByAirportRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlightsByAirport200JSONResponse{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
		{Id: 2, Number: "ST2", OriginAirport: api.Airport{Id: 2}, DestinationAirport: api.Airport{Id: 1}},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func checkAirportIataCodes(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListAirports(context.Background(), api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	airports := resp.(api.ListAirports200JSONResponse)
	if len(airports) != len(want) {
		t.Errorf("got %d airports, want %d", len(airports), len(want))
	}
	for i, airport := range airports {
		if airport.IataCode != want[i] {
			t.Errorf("got %q, want %q", airport.IataCode, want[i])
		}
	}
}
