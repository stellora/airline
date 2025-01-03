package main

import (
	"context"
	"errors"
	"regexp"
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
				AircraftType:       "B77W",
				StartDate:          openapi_types.Date{Time: time.Date(2025, 1, 1, 18, 19, 0, 0, time.UTC)},
				EndDate:            openapi_types.Date{Time: time.Date(2025, 1, 1, 20, 21, 0, 0, time.UTC)},
				DaysOfWeek:         []int{1, 3, 5, 6, 7},
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
// of the week).
func parseDaysOfWeek(str string) (days []int, err error) {
	for _, c := range str {
		if c < '0' || c > '6' {
			return nil, errors.New("invalid day of week")
		}
		days = append(days, int(c-'0'))
	}
	return days, nil
}

var intString = regexp.MustCompile(`^\d+$`)

// isIntString returns true if str is a string of one or more digits. Unlike strconv.Atoi or
// strconv.ParseInt, it does not allow leading '-' or '+' characters.
func isIntString(str string) bool {
	return intString.MatchString(str)
}

func ptrTo[T any](v T) *T {
	return &v
}
