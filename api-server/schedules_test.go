package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetSchedule(ctx, api.GetScheduleRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetSchedule200JSONResponse{
			Id:                 1,
			Airline:            xxAirline,
			Number:             "1",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true,
		}
		assertEqual(t, resp, want)
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetSchedule(ctx, api.GetScheduleRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetSchedule404Response{})
	})
}

func TestListSchedules(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	resp, err := handler.ListSchedules(ctx, api.ListSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListSchedules200JSONResponse{
		api.Schedule{
			Id:                 1,
			Airline:            xxAirline,
			Number:             "1",
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true,
		},
		api.Schedule{
			Id:                 2,
			Airline:            xxAirline,
			Number:             "2",
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			Fleet:              ffFleet,
			StartDate:          fixtureLocalDate1.String(),
			EndDate:            fixtureLocalDate2.String(),
			DaysOfWeek:         fixtureDaysOfWeek,
			DepartureTime:      "07:00",
			DurationSec:        durationSec(2, 0),
			Published:          true,
		},
	}
	assertEqual(t, resp, want)
}

func TestCreateSchedule(t *testing.T) {
	t.Run("with airport IDs", func(t *testing.T) {
		ctx, handler := handlerTest(t)
		insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
		insertAirlinesWithIATACodesT(t, handler, "XX")
		insertTestFleet(t, handler)

		resp, err := handler.CreateSchedule(ctx, api.CreateScheduleRequestObject{
			Body: &api.CreateScheduleJSONRequestBody{
				Airline:            api.NewAirlineSpec(0, "XX"),
				Number:             "1",
				OriginAirport:      api.NewAirportSpec(1, ""),
				DestinationAirport: api.NewAirportSpec(2, ""),
				Fleet:              ffFleetSpec,
				StartDate:          fixtureLocalDate1.String(),
				EndDate:            fixtureLocalDate2.String(),
				DepartureTime:      "07:00",
				DurationSec:        durationSec(2, 0),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateSchedule201JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
		}
		checkScheduleTitles(t, handler, []string{"XX1 AAA-BBB"})
	})

	t.Run("with airport IATA codes", func(t *testing.T) {
		ctx, handler := handlerTest(t)
		insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
		insertAirlinesWithIATACodesT(t, handler, "XX")
		insertTestFleet(t, handler)

		resp, err := handler.CreateSchedule(ctx, api.CreateScheduleRequestObject{
			Body: &api.CreateScheduleJSONRequestBody{
				Airline:            api.NewAirlineSpec(0, "XX"),
				Number:             "1",
				OriginAirport:      api.NewAirportSpec(0, "AAA"),
				DestinationAirport: api.NewAirportSpec(0, "BBB"),
				Fleet:              ffFleetSpec,
				StartDate:          fixtureLocalDate1.String(),
				EndDate:            fixtureLocalDate2.String(),
				DepartureTime:      "07:00",
				DurationSec:        durationSec(2, 0),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.CreateSchedule201JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
		}
		checkScheduleTitles(t, handler, []string{"XX1 AAA-BBB"})
	})
}

func TestUpdateSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")
	insertSchedulesT(t, handler, "XX1 AAA-BBB")
	insertFleetT(t, handler, "XX", "FF2", "")

	{
		// Update the flight
		resp, err := handler.UpdateSchedule(ctx, api.UpdateScheduleRequestObject{
			Id: 1,
			Body: &api.UpdateScheduleJSONRequestBody{
				Number:             ptrTo("100"),
				OriginAirport:      ptrTo(api.NewAirportSpec(2, "")),
				DestinationAirport: ptrTo(api.NewAirportSpec(0, "AAA")),
				Fleet:              ptrTo(api.NewFleetSpec(0, "FF2")),
				StartDate:          ptrTo(fixtureLocalDate3.String()),
				EndDate:            ptrTo(fixtureLocalDate4.String()),
				DaysOfWeek:         ptrTo([]int{1, 5}),
				DepartureTime:      ptrTo("08:15"),
				DurationSec:        ptrTo(durationSec(2, 30)),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.UpdateSchedule200JSONResponse); !ok {
			t.Errorf("got %T, want non-error response", resp)
		}
		checkScheduleTitles(t, handler, []string{"XX100 BBB-AAA"})
	}

	{
		// Verify the flight was actually updated
		resp, err := handler.GetSchedule(ctx, api.GetScheduleRequestObject{Id: 1})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetSchedule200JSONResponse{
			Id:                 1,
			Airline:            xxAirline,
			Number:             "100",
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			Fleet:              api.Fleet{Id: 2, Airline: xxAirline, Code: "FF2"},
			StartDate:          fixtureLocalDate3.String(),
			EndDate:            fixtureLocalDate4.String(),
			DaysOfWeek:         []int{1, 5},
			DepartureTime:      "08:15",
			DurationSec:        durationSec(2, 30),
			Published:          true,
		}
		assertEqual(t, resp, want)
	}
}

func TestDeleteSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB")

	resp, err := handler.DeleteSchedule(ctx, api.DeleteScheduleRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteSchedule204Response{}
	assertEqual(t, resp, want)

	checkScheduleTitles(t, handler, []string{})
}

func TestDeleteAllSchedules(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA")

	resp, err := handler.DeleteAllSchedules(ctx, api.DeleteAllSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllSchedules204Response{}
	assertEqual(t, resp, want)

	checkScheduleTitles(t, handler, []string{})
}

func scheduleTitle(flight api.Schedule) string {
	return fmt.Sprintf("%s%s %s-%s", flight.Airline.IataCode, flight.Number, flight.OriginAirport.IataCode, flight.DestinationAirport.IataCode)
}

func checkScheduleTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListSchedules(context.Background(), api.ListSchedulesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	schedules := resp.(api.ListSchedules200JSONResponse)
	if len(schedules) != len(want) {
		t.Errorf("got %d schedules, want %d", len(schedules), len(want))
	}
	assertEqual(t, mapSlice(scheduleTitle, schedules), want)
}
