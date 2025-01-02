package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlight200JSONResponse{
			Id:                 1,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "1",
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
			Published:          true,
		}
		assertEqual(t, resp, want)
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetFlight404Response{})
	})
}

func TestListFlightSchedules(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	resp, err := handler.ListFlightSchedules(ctx, api.ListFlightSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListFlightSchedules200JSONResponse{
		api.Flight{
			Id:                 1,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "1",
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
			Published:          true,
		},
		api.Flight{
			Id:                 2,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "2",
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
			Published:          true,
		},
	}
	assertEqual(t, resp, want)
}

func TestCreateFlightSchedule(t *testing.T) {
	t.Run("with airport IDs", func(t *testing.T) {
		ctx, handler := handlerTest(t)
		insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
		insertAirlinesWithIATACodesT(t, handler, "XX")

		resp, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{
			Body: &api.CreateFlightScheduleJSONRequestBody{
				Airline:            api.NewAirlineSpec(0, "XX"),
				Number:             "1",
				OriginAirport:      api.NewAirportSpec(1, ""),
				DestinationAirport: api.NewAirportSpec(2, ""),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateFlightSchedule201JSONResponse); !ok {
			t.Errorf("got %#v", resp)
		}
		checkFlightTitles(t, handler, []string{"XX1 AAA-BBB"})
	})

	t.Run("with airport IATA codes", func(t *testing.T) {
		ctx, handler := handlerTest(t)
		insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
		insertAirlinesWithIATACodesT(t, handler, "XX")

		resp, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{
			Body: &api.CreateFlightScheduleJSONRequestBody{
				Airline:            api.NewAirlineSpec(0, "XX"),
				Number:             "1",
				OriginAirport:      api.NewAirportSpec(0, "AAA"),
				DestinationAirport: api.NewAirportSpec(0, "BBB"),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateFlightSchedule201JSONResponse); !ok {
			t.Errorf("got %#v", resp)
		}
		checkFlightTitles(t, handler, []string{"XX1 AAA-BBB"})
	})
}

func TestUpdateFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB")

	{
		// Update the flight
		resp, err := handler.UpdateFlight(ctx, api.UpdateFlightRequestObject{
			Id: 1,
			Body: &api.UpdateFlightJSONRequestBody{
				Airline:            ptrTo(api.NewAirlineSpec(0, "YY")),
				Number:             ptrTo("100"),
				OriginAirport:      ptrTo(api.NewAirportSpec(2, "")),
				DestinationAirport: ptrTo(api.NewAirportSpec(0, "AAA")),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.UpdateFlight200JSONResponse); !ok {
			t.Errorf("got %#v", resp)
		}
		checkFlightTitles(t, handler, []string{"YY100 BBB-AAA"})
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlight(ctx, api.GetFlightRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlight200JSONResponse{
			Id:                 1,
			Airline:            api.Airline{Id: 2, IataCode: "YY"},
			Number:             "100",
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
			Published:          true,
		}
		assertEqual(t, resp, want)
	}
}

func TestDeleteFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB")

	resp, err := handler.DeleteFlightSchedule(ctx, api.DeleteFlightScheduleRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteFlightSchedule204Response{}
	assertEqual(t, resp, want)

	checkFlightTitles(t, handler, []string{})
}

func TestDeleteAllFlightSchedules(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	resp, err := handler.DeleteAllFlightSchedules(ctx, api.DeleteAllFlightSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllFlightSchedules204Response{}
	assertEqual(t, resp, want)

	checkFlightTitles(t, handler, []string{})
}

func checkFlightTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListFlightSchedules(context.Background(), api.ListFlightSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	flights := resp.(api.ListFlightSchedules200JSONResponse)
	if len(flights) != len(want) {
		t.Errorf("got %d flights, want %d", len(flights), len(want))
	}
	for i, flight := range flights {
		title := fmt.Sprintf("%s%s %s-%s", flight.Airline.IataCode, flight.Number, flight.OriginAirport.IataCode, flight.DestinationAirport.IataCode)
		if title != want[i] {
			t.Errorf("got %q, want %q", title, want[i])
		}
	}
}
