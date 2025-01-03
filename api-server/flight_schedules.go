package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBFlightSchedule(a db.FlightSchedulesView) api.FlightSchedule {
	b := api.FlightSchedule{
		Id: int(a.ID),
		Airline: fromDBAirline(db.Airline{
			ID:       a.AirlineID,
			IataCode: a.AirlineIataCode,
			Name:     a.AirlineName,
		}),
		Number: a.Number,
		OriginAirport: fromDBAirport(db.Airport{
			ID:       a.OriginAirportID,
			IataCode: a.OriginAirportIataCode,
			OadbID:   a.OriginAirportOadbID,
		}),
		DestinationAirport: fromDBAirport(db.Airport{
			ID:       a.DestinationAirportID,
			IataCode: a.DestinationAirportIataCode,
			OadbID:   a.DestinationAirportOadbID,
		}),
		AircraftType: fromAircraftTypeCode(a.AircraftType),
		StartDate:    openapi_types.Date{Time: a.StartDate},
		EndDate:      openapi_types.Date{Time: a.EndDate},
		DaysOfWeek:   a.DaysOfWeek,
		Published:    a.Published,
	}
	b.DistanceMiles = distanceMilesBetweenAirports(b.OriginAirport, b.DestinationAirport)
	return b
}

func (h *Handler) GetFlightSchedule(ctx context.Context, request api.GetFlightScheduleRequestObject) (api.GetFlightScheduleResponseObject, error) {
	flight, err := h.queries.GetFlightSchedule(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFlightSchedule404Response{}, nil
		}
	}
	return api.GetFlightSchedule200JSONResponse(fromDBFlightSchedule(flight)), nil
}

func (h *Handler) ListFlightSchedules(ctx context.Context, request api.ListFlightSchedulesRequestObject) (api.ListFlightSchedulesResponseObject, error) {
	flights, err := h.queries.ListFlightSchedules(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListFlightSchedules200JSONResponse(mapSlice(fromDBFlightSchedule, flights)), nil
}

func (h *Handler) CreateFlightSchedule(ctx context.Context, request api.CreateFlightScheduleRequestObject) (api.CreateFlightScheduleResponseObject, error) {
	if request.Body.Number == "" {
		return nil, fmt.Errorf("number must not be empty")
	}

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

	// TODO(sqs): return HTTP 400 errors with error msg
	originAirport, err := getOrCreateAirportBySpec(ctx, tx, queriesTx, request.Body.OriginAirport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("originAirport %q not found", request.Body.OriginAirport)
		}
		return nil, fmt.Errorf("looking up originAirport: %w", err)
	}
	destinationAirport, err := getOrCreateAirportBySpec(ctx, tx, queriesTx, request.Body.DestinationAirport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("destinationAirport %q not found", request.Body.DestinationAirport)
		}
		return nil, fmt.Errorf("looking up destinationAirport: %w", err)
	}

	created, err := queriesTx.CreateFlightSchedule(ctx, db.CreateFlightScheduleParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		Published:            request.Body.Published != nil && *request.Body.Published,
	})
	if err != nil {
		return nil, err
	}

	flight, err := queriesTx.GetFlightSchedule(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateFlightSchedule201JSONResponse(fromDBFlightSchedule(flight)), nil
}

func (h *Handler) UpdateFlightSchedule(ctx context.Context, request api.UpdateFlightScheduleRequestObject) (api.UpdateFlightScheduleResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	params := db.UpdateFlightScheduleParams{
		ID: int64(request.Id),
	}
	if request.Body.Airline != nil {
		airline, err := getAirlineBySpec(ctx, queriesTx, *request.Body.Airline)
		if err != nil {
			return nil, err
		}
		params.AirlineID = sql.NullInt64{Int64: airline.ID, Valid: true}
	}
	if request.Body.Number != nil {
		params.Number = sql.NullString{String: *request.Body.Number, Valid: true}
	}
	if request.Body.OriginAirport != nil {
		originAirport, err := getOrCreateAirportBySpec(ctx, tx, queriesTx, *request.Body.OriginAirport)
		if err != nil {
			return nil, err
		}
		params.OriginAirportID = sql.NullInt64{Int64: originAirport.ID, Valid: true}
	}
	if request.Body.DestinationAirport != nil {
		destinationAirport, err := getOrCreateAirportBySpec(ctx, tx, queriesTx, *request.Body.DestinationAirport)
		if err != nil {
			return nil, err
		}
		params.DestinationAirportID = sql.NullInt64{Int64: destinationAirport.ID, Valid: true}
	}
	if request.Body.Published != nil {
		params.Published = sql.NullBool{Bool: *request.Body.Published, Valid: true}
	}

	if _, err := queriesTx.UpdateFlightSchedule(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlightSchedule404Response{}, nil
		}
		return nil, err
	}

	flight, err := queriesTx.GetFlightSchedule(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlightSchedule404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateFlightSchedule200JSONResponse(fromDBFlightSchedule(flight)), nil
}

func (h *Handler) DeleteFlightSchedule(ctx context.Context, request api.DeleteFlightScheduleRequestObject) (api.DeleteFlightScheduleResponseObject, error) {
	if err := h.queries.DeleteFlightSchedule(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteFlightSchedule204Response{}, nil
}

func (h *Handler) DeleteAllFlightSchedules(ctx context.Context, request api.DeleteAllFlightSchedulesRequestObject) (api.DeleteAllFlightSchedulesResponseObject, error) {
	if err := h.queries.DeleteAllFlightSchedules(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllFlightSchedules204Response{}, nil
}
