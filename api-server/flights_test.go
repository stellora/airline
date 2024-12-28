package main

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

// parseFlightNumberAndRoute parses a string of the form "UA123 SFO-JFK".
func parseFlightNumberAndRoute(str string) (flightNumber, originIATACode, destinationIATACode string) {
	var route string
	flightNumber, route, _ = strings.Cut(str, " ")
	originIATACode, destinationIATACode, _ = strings.Cut(route, "-")
	return
}

func insertFlights(t *testing.T, queries *db.Queries, flightsNumberAndRoute ...string) (ids []int64) {
	t.Helper()
	ctx := context.Background()
	ids = make([]int64, len(flightsNumberAndRoute))
	for i, flight := range flightsNumberAndRoute {
		flightNumber, originIATACode, destinationIATACode := parseFlightNumberAndRoute(flight)

		originAirport, err := queries.GetAirportByIATACode(ctx, originIATACode)
		if err != nil {
			t.Fatal(err)
		}
		destinationAirport, err := queries.GetAirportByIATACode(ctx, destinationIATACode)
		if err != nil {
			t.Fatal(err)
		}

		v, err := queries.CreateFlight(context.Background(), db.CreateFlightParams{
			Number:               flightNumber,
			OriginAirportID:      originAirport.ID,
			DestinationAirportID: destinationAirport.ID,
		})
		if err != nil {
			t.Fatal(err)
		}
		ids[i] = v.ID
	}
	return ids
}

func TestGetFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodes(t, handler.queries, "AAA", "BBB")
	insertFlights(t, handler.queries, "ST1 AAA-BBB", "ST2 BBB-AAA")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlight200JSONResponse{
			Id:                 1,
			Number:             "ST1",
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
			Published:          false,
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 999,
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
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
		{Id: 2, Number: "ST2", OriginAirport: api.Airport{Id: 2}, DestinationAirport: api.Airport{Id: 1}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	resp, err := handler.ListFlights(ctx, api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlights200JSONResponse{
		api.Flight{
			Id:                 1,
			Number:             "ST1",
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
		},
		api.Flight{
			Id:                 2,
			Number:             "ST2",
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
		},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	t.Run("with airport IDs", func(t *testing.T) {
		flights = nil
		resp, err := handler.CreateFlight(ctx, api.CreateFlightRequestObject{
			Body: &api.CreateFlightJSONRequestBody{
				Number:             "ST1",
				OriginAirport:      newAirportSpec(1, ""),
				DestinationAirport: newAirportSpec(2, ""),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.CreateFlight201Response{}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}

		checkFlightTitles(t, handler, []string{"ST1 AAA-BBB"})
	})

	t.Run("with airport IATA codes", func(t *testing.T) {
		flights = nil
		resp, err := handler.CreateFlight(ctx, api.CreateFlightRequestObject{
			Body: &api.CreateFlightJSONRequestBody{
				Number:             "ST1",
				OriginAirport:      newAirportSpec(0, "AAA"),
				DestinationAirport: newAirportSpec(0, "BBB"),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		want := api.CreateFlight201Response{}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}

		checkFlightTitles(t, handler, []string{"ST1 AAA-BBB"})
	})
}

func TestUpdateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	{
		// Update the flight
		resp, err := handler.UpdateFlight(ctx, api.UpdateFlightRequestObject{
			Id: 1,
			Body: &api.UpdateFlightJSONRequestBody{
				Number:             ptrTo("ST100"),
				OriginAirport:      ptrTo(2),
				DestinationAirport: ptrTo(1),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (api.UpdateFlight204Response{}); !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlight200JSONResponse{
			Id:                 1,
			Number:             "ST100",
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
			Published:          true,
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	}
}

func TestDeleteFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	resp, err := handler.DeleteFlight(ctx, api.DeleteFlightRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteFlight204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkFlightNumbers(t, handler, []string{})
}

func TestDeleteAllFlights(t *testing.T) {
	ctx, handler := handlerTest(t)
	flights = []*api.Flight{
		{Id: 1, Number: "ST1", OriginAirport: api.Airport{Id: 1}, DestinationAirport: api.Airport{Id: 2}},
	}
	airports = []*api.Airport{
		{Id: 1, IataCode: "AAA"},
		{Id: 2, IataCode: "BBB"},
	}

	resp, err := handler.DeleteAllFlights(ctx, api.DeleteAllFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllFlights204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkFlightNumbers(t, handler, []string{})
}

func checkFlightNumbers(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListFlights(context.Background(), api.ListFlightsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	flights := resp.(api.ListFlights200JSONResponse)
	if len(flights) != len(want) {
		t.Errorf("got %d flights, want %d", len(flights), len(want))
	}
	for i, flight := range flights {
		if flight.Number != want[i] {
			t.Errorf("got %q, want %q", flight.Number, want[i])
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
	for i, flight := range flights {
		title := fmt.Sprintf("%s %s-%s", flight.Number, flight.OriginAirport.IataCode, flight.DestinationAirport.IataCode)
		if title != want[i] {
			t.Errorf("got %q, want %q", title, want[i])
		}
	}
}
