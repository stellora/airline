package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func getAircraftBySpec(ctx context.Context, queriesTx *db.Queries, spec api.AircraftSpec) (db.AircraftView, error) {
	if id, err := spec.AsAircraftID(); err == nil {
		return queriesTx.GetAircraft(ctx, int64(id))
	}
	if registration, err := spec.AsAircraftRegistration(); err == nil {
		return queriesTx.GetAircraftByRegistration(ctx, registration)
	}
	panic("invalid AircraftSpec: " + fmt.Sprintf("%#v", spec))
}

func fromDBAircraft(a db.AircraftView) api.Aircraft {
	b := api.Aircraft{
		Id:           int(a.ID),
		Registration: a.Registration,
		AircraftType: a.AircraftType,
		Airline: api.Airline{
			Id:       int(a.AirlineID),
			IataCode: a.AirlineIataCode,
			Name:     a.AirlineName,
		},
	}
	return b
}

func (h *Handler) GetAircraft(ctx context.Context, request api.GetAircraftRequestObject) (api.GetAircraftResponseObject, error) {
	aircraft, err := getAircraftBySpec(ctx, h.queries, request.AircraftSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetAircraft404Response{}, nil
		}
	}
	return api.GetAircraft200JSONResponse(fromDBAircraft(aircraft)), nil
}

func (h *Handler) ListAircraft(ctx context.Context, request api.ListAircraftRequestObject) (api.ListAircraftResponseObject, error) {
	aircraft, err := h.queries.ListAircraft(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListAircraft200JSONResponse(mapSlice(fromDBAircraft, aircraft)), nil
}

func (h *Handler) CreateAircraft(ctx context.Context, request api.CreateAircraftRequestObject) (api.CreateAircraftResponseObject, error) {
	// TODO!(sqs): validate aircraft type IATA code

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.Body.Airline)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("airline %q not found", request.Body.Airline)
		}
		return nil, err
	}

	created, err := queriesTx.CreateAircraft(ctx, db.CreateAircraftParams{
		Registration: request.Body.Registration,
		AircraftType: request.Body.AircraftType,
		AirlineID:    airline.ID,
	})
	if err != nil {
		return api.CreateAircraft400Response{}, err
	}

	aircraft, err := queriesTx.GetAircraft(ctx, created.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateAircraft201JSONResponse(fromDBAircraft(aircraft)), nil
}

func (h *Handler) UpdateAircraft(ctx context.Context, request api.UpdateAircraftRequestObject) (api.UpdateAircraftResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	existing, err := getAircraftBySpec(ctx, queriesTx, request.AircraftSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAircraft404Response{}, nil
		}
		return nil, err
	}

	params := db.UpdateAircraftParams{ID: existing.ID}
	if request.Body.Registration != nil {
		params.Registration = sql.NullString{String: *request.Body.Registration, Valid: true}
	}
	if request.Body.AircraftType != nil {
		params.AircraftType = sql.NullString{String: *request.Body.AircraftType, Valid: true}
	}
	if request.Body.Airline != nil {
		airline, err := getAirlineBySpec(ctx, queriesTx, *request.Body.Airline)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("airline %q not found", request.Body.Airline)
			}
			return nil, err
		}
		params.AirlineID = sql.NullInt64{Int64: airline.ID, Valid: true}
	}

	if _, err := queriesTx.UpdateAircraft(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAircraft404Response{}, nil
		}
		return nil, err
	}

	aircraft, err := queriesTx.GetAircraft(ctx, existing.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAircraft404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateAircraft200JSONResponse(fromDBAircraft(aircraft)), nil
}

func (h *Handler) DeleteAircraft(ctx context.Context, request api.DeleteAircraftRequestObject) (api.DeleteAircraftResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	aircraft, err := getAircraftBySpec(ctx, queriesTx, request.AircraftSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeleteAircraft404Response{}, nil
		}
		return nil, err
	}

	if err := queriesTx.DeleteAircraft(ctx, aircraft.ID); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.DeleteAircraft204Response{}, nil
}

func (h *Handler) DeleteAllAircraft(ctx context.Context, request api.DeleteAllAircraftRequestObject) (api.DeleteAllAircraftResponseObject, error) {
	if err := h.queries.DeleteAllAircraft(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllAircraft204Response{}, nil
}
