package main

import (
	"context"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightsByAirport(ctx context.Context, request api.ListFlightsByAirportRequestObject) (api.ListFlightsByAirportResponseObject, error) {
	airportId := request.Id

	flights, err := h.queries.ListFlightsByAirport(ctx, int64(airportId))
	if err != nil {
		return nil, err
	}
	return api.ListFlightsByAirport200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
