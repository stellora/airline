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

func fromDBFlightInstance(a db.FlightInstancesView) api.FlightInstance {
	// TODO!(sqs): bring in aircraft
	b := api.FlightInstance{
		Id:                   int(a.ID),
		ScheduleID:           nullInt64(a.SourceFlightScheduleID),
		ScheduleInstanceDate: nullTime(a.SourceFlightScheduleInstanceDate),
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
		AircraftType:      fromAircraftTypeCode(a.AircraftType),
		DepartureDateTime: openapi_types.Date{Time: a.DepartureDatetime},
		ArrivalDateTime:   openapi_types.Date{Time: a.ArrivalDatetime},
		Notes:             a.Notes,
		Published:         a.Published,
	}
	if a.AircraftID.Valid {
		b.Aircraft = ptrTo(fromDBAircraft(db.AircraftView{
			ID:              a.AircraftID.Int64,
			Registration:    a.AircraftRegistration,
			AircraftType:    a.AircraftAircraftType,
			AirlineID:       a.AircraftAirlineID,
			AirlineIataCode: a.AircraftAirlineIataCode,
			AirlineName:     a.AircraftAirlineName,
		}))
	}
	return b
}

func (h *Handler) GetFlightInstance(ctx context.Context, request api.GetFlightInstanceRequestObject) (api.GetFlightInstanceResponseObject, error) {
	row, err := h.queries.GetFlightInstance(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFlightInstance404Response{}, nil
		}
	}
	return api.GetFlightInstance200JSONResponse(fromDBFlightInstance(row)), nil
}

func (h *Handler) ListFlightInstances(ctx context.Context, request api.ListFlightInstancesRequestObject) (api.ListFlightInstancesResponseObject, error) {
	rows, err := h.queries.ListFlightInstances(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstances200JSONResponse(mapSlice(fromDBFlightInstance, rows)), nil
}

func (h *Handler) CreateFlightInstance(ctx context.Context, request api.CreateFlightInstanceRequestObject) (api.CreateFlightInstanceResponseObject, error) {
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

	if request.Body.AircraftType == "" {
		return nil, errors.New("aircraftType must not be empty")
	}

	var aircraftID sql.NullInt64
	if request.Body.Aircraft != nil {
		aircraft, err := getAircraftBySpec(ctx, queriesTx, *request.Body.Aircraft)
		if err != nil {
			return nil, err
		}
		aircraftID = sql.NullInt64{Valid: true, Int64: aircraft.ID}
	}

	created, err := queriesTx.CreateFlightInstance(ctx, db.CreateFlightInstanceParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		AircraftType:         request.Body.AircraftType,
		AircraftID:           aircraftID,
		DepartureDatetime:    request.Body.DepartureDateTime.Time,
		ArrivalDatetime:      request.Body.ArrivalDateTime.Time,
		Notes:                request.Body.Notes,
		Published:            request.Body.Published != nil && *request.Body.Published,
	})
	if err != nil {
		return nil, err
	}

	row, err := queriesTx.GetFlightInstance(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateFlightInstance201JSONResponse(fromDBFlightInstance(row)), nil
}

func (h *Handler) UpdateFlightInstance(ctx context.Context, request api.UpdateFlightInstanceRequestObject) (api.UpdateFlightInstanceResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	params := db.UpdateFlightInstanceParams{
		ID: int64(request.Id),
	}
	if request.Body.Aircraft != nil {
		aircraft, err := getAircraftBySpec(ctx, queriesTx, *request.Body.Aircraft)
		if err != nil {
			return nil, err
		}
		params.AircraftID = sql.NullInt64{Int64: aircraft.ID, Valid: true}
	}
	if request.Body.Notes != nil {
		params.Notes = sql.NullString{String: *request.Body.Notes, Valid: true}
	}

	if _, err := queriesTx.UpdateFlightInstance(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlightInstance404Response{}, nil
		}
		return nil, err
	}

	row, err := queriesTx.GetFlightInstance(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlightInstance404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateFlightInstance200JSONResponse(fromDBFlightInstance(row)), nil
}

func (h *Handler) DeleteFlightInstance(ctx context.Context, request api.DeleteFlightInstanceRequestObject) (api.DeleteFlightInstanceResponseObject, error) {
	if err := h.queries.DeleteFlightInstance(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteFlightInstance204Response{}, nil
}
