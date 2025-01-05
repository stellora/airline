package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func getPassengerByID(ctx context.Context, queries *db.Queries, id int64) (db.Passenger, error) {
	return queries.GetPassenger(ctx, id)
}

func fromDBPassenger(p db.Passenger) api.Passenger {
	return api.Passenger{
		Id:   int(p.ID),
		Name: p.Name,
	}
}

func (h *Handler) GetPassenger(ctx context.Context, request api.GetPassengerRequestObject) (api.GetPassengerResponseObject, error) {
	passenger, err := getPassengerByID(ctx, h.queries, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetPassenger404Response{}, nil
		}
		return nil, err
	}
	return api.GetPassenger200JSONResponse(fromDBPassenger(passenger)), nil
}

func (h *Handler) ListPassengers(ctx context.Context, request api.ListPassengersRequestObject) (api.ListPassengersResponseObject, error) {
	passengers, err := h.queries.ListPassengers(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListPassengers200JSONResponse(mapSlice(fromDBPassenger, passengers)), nil
}

func (h *Handler) CreatePassenger(ctx context.Context, request api.CreatePassengerRequestObject) (api.CreatePassengerResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	id, err := queriesTx.CreatePassenger(ctx, request.Body.Name)
	if err != nil {
		return nil, err
	}

	created, err := queriesTx.GetPassenger(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreatePassenger201JSONResponse(fromDBPassenger(created)), nil
}

func (h *Handler) UpdatePassenger(ctx context.Context, request api.UpdatePassengerRequestObject) (api.UpdatePassengerResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	passenger, err := getPassengerByID(ctx, queriesTx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdatePassenger404Response{}, nil
		}
		return nil, err
	}

	params := db.UpdatePassengerParams{ID: passenger.ID}
	if request.Body.Name != nil {
		params.Name = sql.NullString{String: *request.Body.Name, Valid: true}
	}

	id, err := queriesTx.UpdatePassenger(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdatePassenger404Response{}, nil
		}
		return nil, err
	}

	updated, err := queriesTx.GetPassenger(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdatePassenger200JSONResponse(fromDBPassenger(updated)), nil
}

func (h *Handler) DeletePassenger(ctx context.Context, request api.DeletePassengerRequestObject) (api.DeletePassengerResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	passenger, err := getPassengerByID(ctx, queriesTx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeletePassenger404Response{}, nil
		}
		return nil, err
	}

	if err := queriesTx.DeletePassenger(ctx, passenger.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.DeletePassenger204Response{}, nil
}
