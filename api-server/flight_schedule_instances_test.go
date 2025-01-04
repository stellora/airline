package main

import (
	"testing"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/zonedtime"
)

func TestSyncFlightInstancesForFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")

	flightSchedule1 := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)

	// Ensure that the flight instances are preserved when updating the schedule.
	setNotesForFlightInstance(t, handler, 1, "a")
	setNotesForFlightInstance(t, handler, 2, "b")
	setNotesForFlightInstance(t, handler, 3, "c")
	checkFlightInstances(t, handler, flightSchedule1.Id, []string{
		"2025-01-01 notes=a",
		"2025-01-02 notes=b",
		"2025-01-03 notes=c",
	})

	_, err := handler.UpdateFlightSchedule(ctx, api.UpdateFlightScheduleRequestObject{
		Id: flightSchedule1.Id,
		Body: &api.UpdateFlightScheduleJSONRequestBody{
			EndDate:    ptrTo("2025-01-05"),
			DaysOfWeek: ptrTo([]int{0, 3, 5, 6}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	checkFlightInstances(t, handler, flightSchedule1.Id, []string{
		"2025-01-01 notes=a",
		"2025-01-03 notes=c",
		"2025-01-04",
		"2025-01-05",
	})
}

func TestListFlightInstancesForFlightSchedule(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-01"),
		mustParseLocalDate("2025-01-03"),
		allDaysOfWeek,
	)
	flightSchedule2 := insertFlightScheduleT(t, handler,
		mustParseLocalDate("2025-01-02"),
		mustParseLocalDate("2025-01-04"),
		allDaysOfWeek,
	)

	resp, err := handler.ListFlightInstancesForFlightSchedule(ctx, api.ListFlightInstancesForFlightScheduleRequestObject{
		Id: flightSchedule2.Id,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListFlightInstancesForFlightSchedule200JSONResponse{
		{
			Id:                   4,
			ScheduleID:           &flightSchedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-02"),
			Airline:              flightSchedule2.Airline,
			Number:               flightSchedule2.Number,
			OriginAirport:        flightSchedule2.OriginAirport,
			DestinationAirport:   flightSchedule2.DestinationAirport,
			AircraftType:         flightSchedule2.AircraftType,
			DepartureDateTime:    zonedtime.ZonedTime{time.Date(2025, 1, 2, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{time.Date(2025, 1, 2, 9, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            flightSchedule2.Published,
		},
		{
			Id:                   5,
			ScheduleID:           &flightSchedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-03"),
			Airline:              flightSchedule2.Airline,
			Number:               flightSchedule2.Number,
			OriginAirport:        flightSchedule2.OriginAirport,
			DestinationAirport:   flightSchedule2.DestinationAirport,
			AircraftType:         flightSchedule2.AircraftType,
			DepartureDateTime:    zonedtime.ZonedTime{time.Date(2025, 1, 3, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{time.Date(2025, 1, 3, 9, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            flightSchedule2.Published,
		},
		{
			Id:                   6,
			ScheduleID:           &flightSchedule2.Id,
			ScheduleInstanceDate: ptrTo("2025-01-04"),
			Airline:              flightSchedule2.Airline,
			Number:               flightSchedule2.Number,
			OriginAirport:        flightSchedule2.OriginAirport,
			DestinationAirport:   flightSchedule2.DestinationAirport,
			AircraftType:         flightSchedule2.AircraftType,
			DepartureDateTime:    zonedtime.ZonedTime{time.Date(2025, 1, 4, 7, 0, 0, 0, mustGetTzLocation(aaaAirport.TimezoneID))},
			ArrivalDateTime:      zonedtime.ZonedTime{time.Date(2025, 1, 4, 9, 0, 0, 0, mustGetTzLocation(bbbAirport.TimezoneID))},
			Published:            flightSchedule2.Published,
		},
	})
}
