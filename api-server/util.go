package main

import (
	"context"
	"database/sql"
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

func insertFlightInstance(ctx context.Context, handler *Handler, newInstance api.CreateFlightInstanceJSONRequestBody) (api.FlightInstance, error) {
	v, err := handler.CreateFlightInstance(ctx, api.CreateFlightInstanceRequestObject{
		Body: &newInstance,
	})
	if err != nil {
		return api.FlightInstance{}, err
	}
	return api.FlightInstance(v.(api.CreateFlightInstance201JSONResponse)), nil
}

func distanceMilesBetweenAirports(a, b api.Airport) float64 {
	if (a.Point != api.Point{} && b.Point != api.Point{}) {
		var distanceMeters float64
		geodesic.WGS84.Inverse(a.Point.Latitude, a.Point.Longitude, b.Point.Latitude, b.Point.Longitude, &distanceMeters, nil, nil)
		const metersPerMile = 0.000621371192237334
		return distanceMeters * metersPerMile
	}
	panic("unable to compute distanceMilesBetweenAiports")
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

func nullInt64(nv sql.NullInt64) *int {
	if nv.Valid {
		v := int(nv.Int64)
		return &v
	}
	return nil
}

func nullTime(nv sql.NullTime) *openapi_types.Date {
	if nv.Valid {
		return &openapi_types.Date{Time: nv.Time}
	}
	return nil
}
