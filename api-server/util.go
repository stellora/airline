package main

import (
	"context"
	"strings"

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

// parseFlightTitle parses a flight title of the form "UA123 SFO-JFK".
func parseFlightTitle(title string) (flightNumber, originIATACode, destinationIATACode string) {
	var route string
	flightNumber, route, _ = strings.Cut(title, " ")
	originIATACode, destinationIATACode, _ = strings.Cut(route, "-")
	return
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

func insertFlights(ctx context.Context, handler *Handler, flightTitles ...string) (ids []int, err error) {
	ids = make([]int, len(flightTitles))
	for i, flight := range flightTitles {
		flightNumber, originIATACode, destinationIATACode := parseFlightTitle(flight)
		v, err := handler.CreateFlight(ctx, api.CreateFlightRequestObject{
			Body: &api.CreateFlightJSONRequestBody{
				Number:             flightNumber,
				OriginAirport:      newAirportSpec(0, originIATACode),
				DestinationAirport: newAirportSpec(0, destinationIATACode),
				Published:          ptrTo(true),
			},
		})
		if err != nil {
			return nil, err
		}
		ids[i] = v.(api.CreateFlight201JSONResponse).Id
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

func ptrTo[T any](v T) *T {
	return &v
}
