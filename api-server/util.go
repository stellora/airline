package main

import (
	"context"
	"errors"
	"sort"
	"strings"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stellora/airline/api-server/api"
	"github.com/tidwall/geodesic"
)

func mapSlice[T any, U any](fn func(T) U, slice []T) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// parseFlightTitle parses a flight schedule title of the form "UA123 SFO-JFK".
func parseFlightTitle(title string) (airlineIATACode, flightNumber, originIATACode, destinationIATACode string) {
	var airlineFlightNumber, route string
	airlineFlightNumber, route, _ = strings.Cut(title, " ")
	airlineIATACode, flightNumber = airlineFlightNumber[:2], airlineFlightNumber[2:]
	originIATACode, destinationIATACode, _ = strings.Cut(route, "-")
	return
}

func insertAircraft(ctx context.Context, handler *Handler, airlineIATACode, aircraftType string, registrations ...string) (ids []int, err error) {
	ids = make([]int, len(registrations))
	for i, reg := range registrations {
		v, err := handler.CreateAircraft(ctx, api.CreateAircraftRequestObject{
			Body: &api.CreateAircraftJSONRequestBody{
				Registration: reg,
				AircraftType: aircraftType,
				Airline:      api.NewAirlineSpec(0, airlineIATACode),
			},
		})
		if err != nil {
			return nil, err
		}
		ids[i] = v.(api.CreateAircraft201JSONResponse).Id
	}
	return ids, nil
}

func insertAirportsWithIATACodes(ctx context.Context, handler *Handler, iataCodes ...string) (ids []int, err error) {
	ids = make([]int, len(iataCodes))
	for i, iataCode := range iataCodes {
		v, err := handler.CreateAirport(ctx, api.CreateAirportRequestObject{
			Body: &api.CreateAirportJSONRequestBody{IataCode: iataCode},
		})
		if err != nil {
			return nil, err
		}
		ids[i] = v.(api.CreateAirport201JSONResponse).Id
	}
	return ids, nil
}

func insertAirlinesWithIATACodes(ctx context.Context, handler *Handler, iataCodes ...string) (ids []int, err error) {
	ids = make([]int, len(iataCodes))
	for i, iataCode := range iataCodes {
		v, err := handler.CreateAirline(ctx, api.CreateAirlineRequestObject{
			Body: &api.CreateAirlineJSONRequestBody{IataCode: iataCode},
		})
		if err != nil {
			return nil, err
		}
		ids[i] = v.(api.CreateAirline201JSONResponse).Id
	}
	return ids, nil
}

func insertAirlines(ctx context.Context, handler *Handler, airlines map[string]string) (ids []int, err error) {
	ids = make([]int, 0, len(airlines))
	for iataCode, name := range airlines {
		v, err := handler.CreateAirline(ctx, api.CreateAirlineRequestObject{
			Body: &api.CreateAirlineJSONRequestBody{IataCode: iataCode, Name: name},
		})
		if err != nil {
			return nil, err
		}
		ids = append(ids, v.(api.CreateAirline201JSONResponse).Id)
	}
	return ids, nil
}

var (
	fixtureDate1      = openapi_types.Date{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)}
	fixtureDate2      = openapi_types.Date{Time: time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC)}
	fixtureDate3      = openapi_types.Date{Time: time.Date(2025, 1, 25, 0, 0, 0, 0, time.UTC)}
	fixtureDate4      = openapi_types.Date{Time: time.Date(2025, 2, 7, 0, 0, 0, 0, time.UTC)}
	fixtureDaysOfWeek = []int{1, 5, 6}
	fixtureB77W       = api.AircraftType{IcaoCode: "B77W", Name: "Boeing 777-300ER"}

	allDaysOfWeek    = []int{0, 1, 2, 3, 4, 5, 6}
	allDaysOfWeekStr = "0123456"
)

func insertFlightSchedules(ctx context.Context, handler *Handler, flightTitles ...string) (ids []int, err error) {
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
				StartDate:          fixtureDate1,
				EndDate:            fixtureDate2,
				DaysOfWeek:         fixtureDaysOfWeek,
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

func insertFlightSchedule(ctx context.Context, handler *Handler, startDate, endDate time.Time, daysOfWeek []int) (api.FlightSchedule, error) {
	v, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{
		Body: &api.CreateFlightScheduleJSONRequestBody{
			Airline:            api.NewAirlineSpec(0, "XX"),
			Number:             "1",
			OriginAirport:      api.NewAirportSpec(0, "AAA"),
			DestinationAirport: api.NewAirportSpec(0, "BBB"),
			AircraftType:       fixtureB77W.IcaoCode,
			StartDate:          openapi_types.Date{Time: startDate},
			EndDate:            openapi_types.Date{Time: endDate},
			DaysOfWeek:         daysOfWeek,
			Published:          ptrTo(true),
		},
	})
	if err != nil {
		return api.FlightSchedule{}, err
	}
	return api.FlightSchedule(v.(api.CreateFlightSchedule201JSONResponse)), nil
}

func distanceMilesBetweenAirports(a, b api.Airport) *float64 {
	if (a.Point != api.Point{} && b.Point != api.Point{}) {
		var distanceMeters float64
		geodesic.WGS84.Inverse(a.Point.Latitude, a.Point.Longitude, b.Point.Latitude, b.Point.Longitude, &distanceMeters, nil, nil)
		const metersPerMile = 0.000621371192237334
		return ptrTo(distanceMeters * metersPerMile)
	}
	return nil
}

// parseDaysOfWeek parses a string like `01356` to a slice with those numbers (representing the days
// of the week). Sunday is 0.
func parseDaysOfWeek(str string) (days []int, err error) {
	seen := make([]bool, 7)
	for _, c := range str {
		if c < '0' || c > '6' {
			return nil, errors.New("invalid day of week")
		}
		day := int(c - '0')
		if seen[day] {
			continue
		}
		seen[day] = true
		days = append(days, day)
	}
	sort.Ints(days)
	return days, nil
}

func toDBDaysOfWeek(days []int) string {
	s := make([]byte, len(days))
	for i, day := range days {
		s[i] = byte('0' + day)
	}
	return string(s)
}

func daysOfWeekContains(daysOfWeek []int, day time.Weekday) bool {
	for _, d := range daysOfWeek {
		if d == int(day) {
			return true
		}
	}
	return false
}

func ptrTo[T any](v T) *T {
	return &v
}
