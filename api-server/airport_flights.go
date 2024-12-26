package main

import (
	"context"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightsByAirport(ctx context.Context, request api.ListFlightsByAirportRequestObject) (api.ListFlightsByAirportResponseObject, error) {
	airportId := request.Id

	var flightsByAirport []api.Flight
	for _, flight := range flights {
		if flight.OriginAirport.Id == airportId || flight.DestinationAirport.Id == airportId {
			flightsByAirport = append(flightsByAirport, *flight)
		}
	}
	return api.ListFlightsByAirport200JSONResponse(flightsByAirport), nil
}
