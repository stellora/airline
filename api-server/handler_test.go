package main

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
	"github.com/stellora/airline/api-server/localtime"
)

func init() {
	testingNoDistanceCalculations = true
	extdata.Airports = extdata.AirportsDataset{
		Airports: []extdata.Airport{
			{
				ID:         1,
				IATACode:   "AAA",
				TimezoneID: "America/Los_Angeles",
				Name:       "AAA Airport",
				ISOCountry: "US",
				ISORegion:  "US-CA",
			},
			{
				ID:         2,
				IATACode:   "BBB",
				TimezoneID: "America/New_York",
				Name:       "BBB Airport",
				ISOCountry: "US",
				ISORegion:  "US-NY",
			},
			{
				ID:         3,
				IATACode:   "CCC",
				TimezoneID: "America/Chicago",
				Name:       "CCC Airport",
				ISOCountry: "US",
				ISORegion:  "US-IL",
			},
		},
		Regions: map[extdata.ISORegion]extdata.Region{
			"US-CA": {
				Code:       "US-CA",
				Name:       "California",
				ISOCountry: "US",
			},
			"US-NY": {
				Code:       "US-NY",
				Name:       "New York",
				ISOCountry: "US",
			},
			"US-IL": {
				Code:       "US-IL",
				Name:       "Illinois",
				ISOCountry: "US",
			},
		},
		Countries: map[extdata.ISOCountry]extdata.Country{
			"US": {
				Code: "US",
				Name: "United States",
			},
		},
	}
}

func mustGetTzLocation(tzID string) *time.Location {
	loc, err := time.LoadLocation(tzID)
	if err != nil {
		panic(err)
	}
	return loc
}

func mustParseLocalDate(s string) localtime.LocalDate {
	d, err := localtime.ParseLocalDate(s)
	if err != nil {
		panic(err)
	}
	return d
}

func durationSec(hours, minutes int) int {
	return (hours*60 + minutes) * 60
}

var (
	fixtureLocalDate1 = mustParseLocalDate("2025-01-01")
	fixtureLocalDate2 = mustParseLocalDate("2025-01-20")
	fixtureLocalDate3 = mustParseLocalDate("2025-01-25")
	fixtureLocalDate4 = mustParseLocalDate("2025-02-07")
	fixtureDaysOfWeek = []int{1, 5, 6}
	fixtureB77W       = api.AircraftType{IcaoCode: "B77W", Name: "Boeing 777-300ER"}

	allDaysOfWeek = []int{0, 1, 2, 3, 4, 5, 6}

	aaaAirport = api.Airport{
		Id:         1,
		IataCode:   "AAA",
		Name:       "AAA Airport",
		Region:     "California",
		Country:    "US",
		TimezoneID: "America/Los_Angeles",
	}
	bbbAirport = api.Airport{
		Id:         2,
		IataCode:   "BBB",
		Name:       "BBB Airport",
		Region:     "New York",
		Country:    "US",
		TimezoneID: "America/New_York",
	}
	cccAirport = api.Airport{
		Id:         3,
		IataCode:   "CCC",
		Name:       "CCC Airport",
		Region:     "Illinois",
		Country:    "US",
		TimezoneID: "America/Chicago",
	}
)

func handlerTest(t *testing.T) (context.Context, *Handler) {
	ctx := context.Background()
	db, queries, err := db.Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	return ctx, NewHandler(db, queries)
}

func insertAircraftT(t *testing.T, handler *Handler, airline string, registrations ...string) (ids []int) {
	t.Helper()
	ids, err := insertAircraft(context.Background(), handler, airline, "B77W", registrations...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertAirportsWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirportsWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertAirlinesWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirlinesWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightSchedulesT(t *testing.T, handler *Handler, flightTitles ...string) (ids []int) {
	insertFlightSchedules := func(ctx context.Context, handler *Handler, flightTitles ...string) (ids []int, err error) {
		ids = make([]int, len(flightTitles))
		for i, flight := range flightTitles {
			airlineIATACode, flightNumber, originIATACode, destinationIATACode := parseFlightTitle(flight)
			v, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{
				Body: &api.CreateFlightScheduleJSONRequestBody{
					Airline:            api.NewAirlineSpec(0, airlineIATACode),
					Number:             flightNumber,
					OriginAirport:      api.NewAirportSpec(0, originIATACode),
					DestinationAirport: api.NewAirportSpec(0, destinationIATACode),
					AircraftType:       fixtureB77W.IcaoCode,
					StartDate:          fixtureLocalDate1.String(),
					EndDate:            fixtureLocalDate2.String(),
					DaysOfWeek:         fixtureDaysOfWeek,
					DepartureTime:      "07:00",
					DurationSec:        durationSec(2, 0),
					Published:          ptrTo(true),
				},
			})
			if err != nil {
				return nil, err
			}
			ids[i] = v.(api.CreateFlightSchedule201JSONResponse).Id
		}
		return ids, nil
	}

	t.Helper()
	ids, err := insertFlightSchedules(context.Background(), handler, flightTitles...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightScheduleT(t *testing.T, handler *Handler, startDate, endDate localtime.LocalDate, daysOfWeek []int) api.FlightSchedule {
	insertFlightSchedule := func(ctx context.Context, handler *Handler, startDate, endDate localtime.LocalDate, daysOfWeek []int) (api.FlightSchedule, error) {
		v, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{
			Body: &api.CreateFlightScheduleJSONRequestBody{
				Airline:            api.NewAirlineSpec(0, "XX"),
				Number:             "1",
				OriginAirport:      api.NewAirportSpec(0, "AAA"),
				DestinationAirport: api.NewAirportSpec(0, "BBB"),
				AircraftType:       fixtureB77W.IcaoCode,
				StartDate:          startDate.String(),
				EndDate:            endDate.String(),
				DaysOfWeek:         daysOfWeek,
				DepartureTime:      "07:00",
				DurationSec:        durationSec(2, 0),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			return api.FlightSchedule{}, err
		}
		return api.FlightSchedule(v.(api.CreateFlightSchedule201JSONResponse)), nil
	}

	t.Helper()
	flightSchedule, err := insertFlightSchedule(context.Background(), handler, startDate, endDate, daysOfWeek)
	if err != nil {
		t.Fatal(err)
	}
	return flightSchedule
}

func insertFlightInstanceT(t *testing.T, handler *Handler, newInstance api.CreateFlightInstanceJSONRequestBody) api.FlightInstance {
	t.Helper()
	instance, err := insertFlightInstance(context.Background(), handler, newInstance)
	if err != nil {
		t.Fatal(err)
	}
	return instance
}

func setNotesForFlightInstance(t *testing.T, handler *Handler, id int, notes string) {
	t.Helper()
	if _, err := handler.UpdateFlightInstance(context.Background(), api.UpdateFlightInstanceRequestObject{
		Id: id,
		Body: &api.UpdateFlightInstanceJSONRequestBody{
			Notes: ptrTo(notes),
		},
	}); err != nil {
		t.Fatal(err)
	}
}

func assertEqual(t *testing.T, got any, want any) {
	t.Helper()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
