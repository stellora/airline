package main

import (
	"testing"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/zonedtime"
)

func TestSyncFlightsForSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	schedule1 := insertScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)

	// Ensure that the flights are preserved when updating the schedule.
	setNotesForFlight(t, handler, 1, "a")
	setNotesForFlight(t, handler, 2, "b")
	setNotesForFlight(t, handler, 3, "c")
	checkFlights(t, handler, schedule1.Id, []string{
		"XX1 AAA-BBB on 2025-01-01 notes=a",
		"XX1 AAA-BBB on 2025-01-02 notes=b",
		"XX1 AAA-BBB on 2025-01-03 notes=c",
	})

	_, err := handler.UpdateSchedule(ctx, api.UpdateScheduleRequestObject{
		Id: schedule1.Id,
		Body: &api.UpdateScheduleJSONRequestBody{
			EndDate:    ptrTo("2025-01-05"),
			DaysOfWeek: ptrTo([]int{0, 3, 5, 6}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	checkFlights(t, handler, schedule1.Id, []string{
		"XX1 AAA-BBB on 2025-01-01 notes=a",
		"XX1 AAA-BBB on 2025-01-03 notes=c",
		"XX1 AAA-BBB on 2025-01-04",
		"XX1 AAA-BBB on 2025-01-05",
	})
}

func TestListFlightsForSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)
	schedule2 := insertScheduleT(t, handler,
		mustParseLocalDate("2025-01-02"),
		mustParseLocalDate("2025-01-04"),
		allDaysOfWeek,
	)

	resp, err := handler.ListFlightsForSchedule(ctx, api.ListFlightsForScheduleRequestObject{
		Id: schedule2.Id,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListFlightsForSchedule200JSONResponse{
		{
			Id:                   4,
			ScheduleID:           &schedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-02"),
			Airline:              schedule2.Airline,
			Number:               schedule2.Number,
			OriginAirport:        schedule2.OriginAirport,
			DestinationAirport:   schedule2.DestinationAirport,
			Fleet:                schedule2.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    zonedtime.ZonedTime{Time: time.Date(2025, 1, 2, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{Time: time.Date(2025, 1, 2, 12, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            schedule2.Published,
		},
		{
			Id:                   5,
			ScheduleID:           &schedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-03"),
			Airline:              schedule2.Airline,
			Number:               schedule2.Number,
			OriginAirport:        schedule2.OriginAirport,
			DestinationAirport:   schedule2.DestinationAirport,
			Fleet:                schedule2.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    zonedtime.ZonedTime{Time: time.Date(2025, 1, 3, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{Time: time.Date(2025, 1, 3, 12, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            schedule2.Published,
		},
		{
			Id:                   6,
			ScheduleID:           &schedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-04"),
			Airline:              schedule2.Airline,
			Number:               schedule2.Number,
			OriginAirport:        schedule2.OriginAirport,
			DestinationAirport:   schedule2.DestinationAirport,
			Fleet:                schedule2.Fleet,
			Aircraft:             nil,
			DepartureDateTime:    zonedtime.ZonedTime{Time: time.Date(2025, 1, 4, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{Time: time.Date(2025, 1, 4, 12, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            schedule2.Published,
		},
	})
}
