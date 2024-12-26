package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []api.Airport{
		{Id: "A", Title: "Airport A"},
		{Id: "B", Title: "Airport B"},
	}

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: "A",
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetAirport200JSONResponse{
			Id:    "A",
			Title: "Airport A",
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: "999",
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
	airports = []api.Airport{
		{Title: "Airport 1"},
		{Title: "Airport 2"},
	}

	resp, err := handler.ListAirports(ctx, api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAirports200JSONResponse{
		api.Airport{Title: "Airport 1"},
		api.Airport{Title: "Airport 2"},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []api.Airport{}

	resp, err := handler.CreateAirport(ctx, api.CreateAirportRequestObject{
		Body: &api.CreateAirportJSONRequestBody{
			Title: "New Airport",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateAirport201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportTitles(t, handler, []string{"New Airport"})
}

func TestDeleteAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []api.Airport{
		{Id: "1", Title: "Airport 1"},
		{Id: "2", Title: "Airport 2"},
	}

	resp, err := handler.DeleteAirport(ctx, api.DeleteAirportRequestObject{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirport204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportTitles(t, handler, []string{"Airport 2"})
}

func checkAirportTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListAirports(context.Background(), api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	airports := resp.(api.ListAirports200JSONResponse)
	if len(airports) != len(want) {
		t.Errorf("got %d airports, want %d", len(airports), len(want))
	}
	for i, airport := range airports {
		if airport.Title != want[i] {
			t.Errorf("got title %q, want %q", airport.Title, want[i])
		}
	}
}
