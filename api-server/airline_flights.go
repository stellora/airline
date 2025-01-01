package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightsByAirline(ctx context.Context, request api.ListFlightsByAirlineRequestObject) (api.ListFlightsByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}

	airline, err := getAirlineBySpec(ctx, h.queries.WithTx(tx), request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightsByAirline404Response{}, nil
		}
		return nil, err
	}

	flights, err := h.queries.ListFlightsByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightsByAirline200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
