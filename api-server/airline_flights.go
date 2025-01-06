package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) ListFlightSchedulesByAirline(ctx context.Context, request api.ListFlightSchedulesByAirlineRequestObject) (api.ListFlightSchedulesByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightSchedulesByAirline404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightSchedulesByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightSchedulesByAirline200JSONResponse(mapSlice(fromDBFlightSchedule, flights)), nil
}

func (h *Handler) ListFlightInstancesByAirline(ctx context.Context, request api.ListFlightInstancesByAirlineRequestObject) (api.ListFlightInstancesByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightInstancesByAirline404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightInstancesByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstancesByAirline200JSONResponse(mapSlice(fromDBFlightInstance, flights)), nil
}
