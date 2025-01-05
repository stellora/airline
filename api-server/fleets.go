package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func getFleetBySpec(ctx context.Context, queriesTx *db.Queries, airlineID int64, spec api.FleetSpec) (db.FleetsView, error) {
	if id, err := spec.AsFleetID(); err == nil {
		return queriesTx.GetFleet(ctx, int64(id))
	}
	if code, err := spec.AsFleetCode(); err == nil {
		return queriesTx.GetFleetByCode(ctx, db.GetFleetByCodeParams{
			AirlineID: airlineID,
			Code:      code,
		})
	}
	panic("invalid FleetSpec")
}

func fromDBFleet(fleet db.FleetsView) api.Fleet {
	return api.Fleet{
		Id: int(fleet.ID),
		Airline: api.Airline{
			Id:       int(fleet.AirlineID),
			IataCode: fleet.AirlineIataCode,
			Name:     fleet.AirlineName,
		},
		Code:        fleet.Code,
		Description: fleet.Description,
	}
}

func (h *Handler) GetFleet(ctx context.Context, request api.GetFleetRequestObject) (api.GetFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFleet404Response{}, nil
		}
		return nil, err
	}

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFleet404Response{}, nil
		}
		return nil, err
	}

	if fleet.AirlineID != airline.ID {
		return &api.GetFleet404Response{}, nil
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.GetFleet200JSONResponse(fromDBFleet(fleet)), nil
}

func (h *Handler) ListFleetsByAirline(ctx context.Context, request api.ListFleetsByAirlineRequestObject) (api.ListFleetsByAirlineResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFleetsByAirline404Response{}, nil
		}
		return nil, err
	}

	fleets, err := queriesTx.ListFleetsByAirline(ctx, airline.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListFleetsByAirline200JSONResponse(mapSlice(fromDBFleet, fleets)), nil
}

func (h *Handler) CreateFleet(ctx context.Context, request api.CreateFleetRequestObject) (api.CreateFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.CreateFleet404Response{}, nil
		}
		return nil, err
	}

	created, err := queriesTx.CreateFleet(ctx, db.CreateFleetParams{
		AirlineID:   airline.ID,
		Code:        request.Body.Code,
		Description: request.Body.Description,
	})
	if err != nil {
		return nil, err
	}

	fleet, err := queriesTx.GetFleet(ctx, created.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateFleet201JSONResponse(fromDBFleet(fleet)), nil
}

func (h *Handler) UpdateFleet(ctx context.Context, request api.UpdateFleetRequestObject) (api.UpdateFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFleet404Response{}, nil
		}
		return nil, err
	}

	existing, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFleet404Response{}, nil
		}
		return nil, err
	}

	if existing.AirlineID != airline.ID {
		return &api.UpdateFleet404Response{}, nil
	}

	params := db.UpdateFleetParams{ID: existing.ID}
	if request.Body.Code != nil {
		params.Code = sql.NullString{String: *request.Body.Code, Valid: true}
	}
	if request.Body.Description != nil {
		params.Description = sql.NullString{String: *request.Body.Description, Valid: true}
	}

	if _, err := queriesTx.UpdateFleet(ctx, params); err != nil {
		return nil, err
	}

	fleet, err := queriesTx.GetFleet(ctx, existing.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateFleet200JSONResponse(fromDBFleet(fleet)), nil
}

func (h *Handler) DeleteFleet(ctx context.Context, request api.DeleteFleetRequestObject) (api.DeleteFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeleteFleet404Response{}, nil
		}
		return nil, err
	}

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeleteFleet404Response{}, nil
		}
		return nil, err
	}

	if fleet.AirlineID != airline.ID {
		return &api.DeleteFleet404Response{}, nil
	}

	if err := queriesTx.DeleteFleet(ctx, fleet.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.DeleteFleet204Response{}, nil
}

func (h *Handler) ListAircraftByFleet(ctx context.Context, request api.ListAircraftByFleetRequestObject) (api.ListAircraftByFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListAircraftByFleet404Response{}, nil
		}
		return nil, err
	}

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListAircraftByFleet404Response{}, nil
		}
		return nil, err
	}

	if fleet.AirlineID != airline.ID {
		return &api.ListAircraftByFleet404Response{}, nil
	}

	aircraft, err := queriesTx.ListAircraftByFleet(ctx, fleet.ID)
	if err != nil {
		return nil, err
	}
	return api.ListAircraftByFleet200JSONResponse(mapSlice(fromDBAircraft, aircraft)), nil
}

func (h *Handler) AddAircraftToFleet(ctx context.Context, request api.AddAircraftToFleetRequestObject) (api.AddAircraftToFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.AddAircraftToFleet404Response{}, nil
		}
		return nil, err
	}

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.AddAircraftToFleet404Response{}, nil
		}
		return nil, err
	}

	if fleet.AirlineID != airline.ID {
		return &api.AddAircraftToFleet404Response{}, nil
	}

	aircraft, err := getAircraftBySpec(ctx, queriesTx, request.AircraftSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.AddAircraftToFleet404Response{}, nil
		}
		return nil, err
	}

	if aircraft.AirlineID != airline.ID {
		return &api.AddAircraftToFleet400Response{}, nil
	}

	err = queriesTx.AddAircraftToFleet(ctx, db.AddAircraftToFleetParams{
		FleetID:    fleet.ID,
		AircraftID: aircraft.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.AddAircraftToFleet200Response{}, nil
}

func (h *Handler) RemoveAircraftFromFleet(ctx context.Context, request api.RemoveAircraftFromFleetRequestObject) (api.RemoveAircraftFromFleetResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airline, err := getAirlineBySpec(ctx, queriesTx, request.AirlineSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.RemoveAircraftFromFleet404Response{}, nil
		}
		return nil, err
	}

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.FleetSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.RemoveAircraftFromFleet404Response{}, nil
		}
		return nil, err
	}

	if fleet.AirlineID != airline.ID {
		return &api.RemoveAircraftFromFleet404Response{}, nil
	}

	aircraft, err := getAircraftBySpec(ctx, queriesTx, request.AircraftSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.RemoveAircraftFromFleet404Response{}, nil
		}
		return nil, err
	}

	if aircraft.AirlineID != airline.ID {
		return &api.RemoveAircraftFromFleet400Response{}, nil
	}

	err = queriesTx.RemoveAircraftFromFleet(ctx, db.RemoveAircraftFromFleetParams{
		FleetID:    fleet.ID,
		AircraftID: aircraft.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.RemoveAircraftFromFleet204Response{}, nil
}
