package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListSchedulesByAirport(ctx context.Context, request api.ListSchedulesByAirportRequestObject) (api.ListSchedulesByAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airport, err := getAirportBySpec(ctx, queriesTx, request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListSchedulesByAirport404Response{}, nil
		}
		return nil, err
	}

	schedules, err := queriesTx.ListSchedulesByAirport(ctx, airport.ID)
	if err != nil {
		return nil, err
	}
	return api.ListSchedulesByAirport200JSONResponse(mapSlice(fromDBSchedule, schedules)), nil
}
