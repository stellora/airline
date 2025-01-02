package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListAircraftByAirline(ctx context.Context, request api.ListAircraftByAirlineRequestObject) (api.ListAircraftByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListAircraftByAirline404Response{}, nil
		}
		return nil, err
	}

	aircraft, err := queriesTx.ListAircraftByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListAircraftByAirline200JSONResponse(mapSlice(fromDBAircraft, aircraft)), nil
}
