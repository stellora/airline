package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetFlightSchedule(ctx, api.GetFlightScheduleRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlightSchedule200JSONResponse{
			Id:                 1,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "1",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			AircraftType:       fixtureB77W,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "7:00",
			ArrivalTime:        "9:00",
			Published:          true,
		}
		assertEqual(t, resp, want)
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFlightSchedule(ctx, api.GetFlightScheduleRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetFlightSchedule404Response{})
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
		api.FlightSchedule{
			Id:                 1,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "1",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			AircraftType:       fixtureB77W,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "7:00",
			ArrivalTime:        "9:00",
			Published:          true,
		},
		api.FlightSchedule{
			Id:                 2,
			Airline:            api.Airline{Id: 1, IataCode: "XX"},
			Number:             "2",
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			AircraftType:       fixtureB77W,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "7:00",
			ArrivalTime:        "9:00",
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
				AircraftType:       "B77W",
				DepartureTime:      "7:00",
				ArrivalTime:        "9:00",
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateFlightSchedule201JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
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
				AircraftType:       "B77W",
				DepartureTime:      "7:00",
				ArrivalTime:        "9:00",
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateFlightSchedule201JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
		}
		checkFlightTitles(t, handler, []string{"XX1 AAA-BBB"})
	})
}

func TestUpdateFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB")

	{
		// Update the flight
		resp, err := handler.UpdateFlightSchedule(ctx, api.UpdateFlightScheduleRequestObject{
			Id: 1,
			Body: &api.UpdateFlightScheduleJSONRequestBody{
				Airline:            ptrTo(api.NewAirlineSpec(0, "YY")),
				Number:             ptrTo("100"),
				OriginAirport:      ptrTo(api.NewAirportSpec(2, "")),
				DestinationAirport: ptrTo(api.NewAirportSpec(0, "AAA")),
				AircraftType:       ptrTo("B789"),
				StartDate:          ptrTo(fixtureLocalDate3.String()),
				EndDate:            ptrTo(fixtureLocalDate4.String()),
				DaysOfWeek:         ptrTo([]int{1, 5}),
				DepartureTime:      ptrTo("8:15"),
				ArrivalTime:        ptrTo("10:45"),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.UpdateFlightSchedule200JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
		}
		checkFlightTitles(t, handler, []string{"YY100 BBB-AAA"})
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetFlightSchedule(ctx, api.GetFlightScheduleRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetFlightSchedule200JSONResponse{
			Id:                 1,
			Airline:            api.Airline{Id: 2, IataCode: "YY"},
			Number:             "100",
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			AircraftType:       api.AircraftType{IcaoCode: "B789", Name: "Boeing 787-9"},
			StartDate:          fixtureLocalDate3.String(),
			EndDate:            fixtureLocalDate4.String(),
			DaysOfWeek:         []int{1, 5},
			DepartureTime:      "8:15",
			ArrivalTime:        "10:45",
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
