package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightSchedulesByAirport(ctx context.Context, request api.ListFlightSchedulesByAirportRequestObject) (api.ListFlightSchedulesByAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airport, err := getAirportBySpec(ctx, queriesTx, request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightSchedulesByAirport404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightSchedulesByAirport(ctx, airport.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightSchedulesByAirport200JSONResponse(mapSlice(fromDBFlightSchedule, flights)), nil
}
