package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightsByAirport(ctx context.Context, request api.ListFlightsByAirportRequestObject) (api.ListFlightsByAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airport, err := getAirportBySpec(ctx, queriesTx, request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightsByAirport404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightsByAirport(ctx, airport.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightsByAirport200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
