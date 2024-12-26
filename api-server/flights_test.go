package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestDeleteAllFlights(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Title: "Flight 1"},
		{Title: "Flight 2"},
	}

	resp, err := handler.DeleteAllFlights(ctx, api.DeleteAllFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllFlights204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkFlightTitles(t, handler, []string{})
}

func TestGetFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Id: "1", Title: "Flight 1"},
		{Id: "2", Title: "Flight 2"},
	}
	airports = []api.Airport{
		{Id: "A", Title: "Airport A"},
	}
	flightAirportMemberships = []flightAirportMembership{
		{flight: "1", airport: "A"},
	}

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: "1",
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlight200JSONResponse{
			Id:         "1",
			Title:      "Flight 1",
			Airports: &[]api.Airport{{Id: "A", Title: "Airport A"}},
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: "999",
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetFlight404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})
}

func TestListFlights(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Id: "1", Title: "Flight 1"},
		{Id: "2", Title: "Flight 2"},
	}
	airports = []api.Airport{
		{Id: "A", Title: "Airport A"},
	}
	flightAirportMemberships = []flightAirportMembership{
		{flight: "1", airport: "A"},
	}

	resp, err := handler.ListFlights(ctx, api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlights200JSONResponse{
		api.Flight{
			Id:         "1",
			Title:      "Flight 1",
			Airports: &[]api.Airport{{Id: "A", Title: "Airport A"}},
		},
		api.Flight{Id: "2", Title: "Flight 2", Airports: &[]api.Airport{}},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestListFlightsByAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Id: "1", Title: "Flight 1"},
		{Id: "2", Title: "Flight 2"},
		{Id: "3", Title: "Flight 3"},
	}
	airports = []api.Airport{
		{Id: "A", Title: "Airport A"},
	}
	flightAirportMemberships = []flightAirportMembership{
		{flight: "1", airport: "A"},
		{flight: "2", airport: "A"},
	}

	resp, err := handler.ListFlightsByAirport(ctx, api.ListFlightsByAirportRequestObject{
		AirportId: "A",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlightsByAirport200JSONResponse{
		FlightsInAirport:    []api.Flight{{Id: "1", Title: "Flight 1"}, {Id: "2", Title: "Flight 2"}},
		FlightsNotInAirport: []api.Flight{{Id: "3", Title: "Flight 3"}},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{}

	resp, err := handler.CreateFlight(ctx, api.CreateFlightRequestObject{
		Body: &api.CreateFlightJSONRequestBody{
			Title: "New Flight",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateFlight201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkFlightTitles(t, handler, []string{"New Flight"})
}

func TestDeleteFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Id: "1", Title: "Flight 1"},
		{Id: "2", Title: "Flight 2"},
	}

	resp, err := handler.DeleteFlight(ctx, api.DeleteFlightRequestObject{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteFlight204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkFlightTitles(t, handler, []string{"Flight 2"})
}

func TestSetFlightStarred(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []api.Flight{
		{Id: "1", Title: "Flight 1", Starred: false},
		{Id: "2", Title: "Flight 2", Starred: false},
	}

	resp, err := handler.SetFlightStarred(ctx, api.SetFlightStarredRequestObject{
		Id: "1",
		Body: &api.SetFlightStarredJSONRequestBody{
			Starred: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.SetFlightStarred204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	// Verify the flight was actually starred
	listResp, err := handler.ListFlights(ctx, api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	flights := listResp.(api.ListFlights200JSONResponse)
	for _, p := range flights {
		if p.Id == "1" && !p.Starred {
			t.Error("Flight 1 should be starred")
		}
		if p.Id == "2" && p.Starred {
			t.Error("Flight 2 should not be starred")
		}
	}
}
func checkFlightTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListFlights(context.Background(), api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	flights := resp.(api.ListFlights200JSONResponse)
	if len(flights) != len(want) {
		t.Errorf("got %d flights, want %d", len(flights), len(want))
	}
	for i, airport := range flights {
		if airport.Title != want[i] {
			t.Errorf("got title %q, want %q", airport.Title, want[i])
		}
	}
}
