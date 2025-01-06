package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListSchedulesByAirline(ctx context.Context, request api.ListSchedulesByAirlineRequestObject) (api.ListSchedulesByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListSchedulesByAirline404Response{}, nil
		}
		return nil, err
	}

	schedules, err := queriesTx.ListSchedulesByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListSchedulesByAirline200JSONResponse(mapSlice(fromDBSchedule, schedules)), nil
}

func (h *Handler) ListFlightsByAirline(ctx context.Context, request api.ListFlightsByAirlineRequestObject) (api.ListFlightsByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightsByAirline404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightsByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightsByAirline200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
