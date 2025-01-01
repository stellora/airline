package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightsByAirport(ctx context.Context, request api.ListFlightsByAirportRequestObject) (api.ListFlightsByAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}

	airport, err := getAirportBySpec(ctx, h.queries.WithTx(tx), request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightsByAirport404Response{}, nil
		}
		return nil, err
	}

	flights, err := h.queries.ListFlightsByAirport(ctx, airport.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightsByAirport200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
