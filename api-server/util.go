package main

import (
	"context"
	"strings"

	"github.com/stellora/airline/api-server/db"
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

func insertAirportsWithIATACodes(ctx context.Context, queries *db.Queries, iataCodes ...string) (ids []int64, err error) {
	ids = make([]int64, len(iataCodes))
	for i, iataCode := range iataCodes {
		v, err := queries.CreateAirport(context.Background(), db.CreateAirportParams{IataCode: iataCode})
		if err != nil {
			return nil, err
		}
		ids[i] = v.ID
	}
	return ids, nil
}

func insertFlights(ctx context.Context, queries *db.Queries, flightTitles ...string) (ids []int64, err error) {
	ids = make([]int64, len(flightTitles))
	for i, flight := range flightTitles {
		flightNumber, originIATACode, destinationIATACode := parseFlightTitle(flight)

		originAirport, err := queries.GetAirportByIATACode(ctx, originIATACode)
		if err != nil {
			return nil, err
		}
		destinationAirport, err := queries.GetAirportByIATACode(ctx, destinationIATACode)
		if err != nil {
			return nil, err
		}

		v, err := queries.CreateFlight(ctx, db.CreateFlightParams{
			Number:               flightNumber,
			OriginAirportID:      originAirport.ID,
			DestinationAirportID: destinationAirport.ID,
		})
		if err != nil {
			return nil, err
		}
		ids[i] = v.ID
	}
	return ids, nil
}
