package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBFlight(a db.FlightsView) api.Flight {
	b := api.Flight{
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
		Published: a.Published,
	}
	b.DistanceMiles = distanceMilesBetweenAirports(b.OriginAirport, b.DestinationAirport)
	return b
}

func (h *Handler) GetFlight(ctx context.Context, request api.GetFlightRequestObject) (api.GetFlightResponseObject, error) {
	flight, err := h.queries.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFlight404Response{}, nil
		}
	}
	return api.GetFlight200JSONResponse(fromDBFlight(flight)), nil
}

func (h *Handler) ListFlights(ctx context.Context, request api.ListFlightsRequestObject) (api.ListFlightsResponseObject, error) {
	flights, err := h.queries.ListFlights(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListFlights200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}

func (h *Handler) CreateFlight(ctx context.Context, request api.CreateFlightRequestObject) (api.CreateFlightResponseObject, error) {
	if request.Body.Number == "" {
		return nil, fmt.Errorf("number must not be empty")
	}

	airline, err := getAirlineBySpec(ctx, h.queries, request.Body.Airline)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("airline %q not found", request.Body.Airline)
		}
		return nil, err
	}

	// TODO(sqs): return HTTP 400 errors with error msg
	originAirport, err := getAirportBySpec(ctx, h.queries, request.Body.OriginAirport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("originAirport %q not found", request.Body.OriginAirport)
		}
		return nil, fmt.Errorf("looking up originAirport: %w", err)
	}
	destinationAirport, err := getAirportBySpec(ctx, h.queries, request.Body.DestinationAirport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("destinationAirport %q not found", request.Body.DestinationAirport)
		}
		return nil, fmt.Errorf("looking up destinationAirport: %w", err)
	}

	created, err := h.queries.CreateFlight(ctx, db.CreateFlightParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		Published:            request.Body.Published != nil && *request.Body.Published,
	})
	if err != nil {
		return nil, err
	}

	flight, err := h.queries.GetFlight(ctx, created)
	if err != nil {
		return nil, err
	}
	return api.CreateFlight201JSONResponse(fromDBFlight(flight)), nil
}

func (h *Handler) UpdateFlight(ctx context.Context, request api.UpdateFlightRequestObject) (api.UpdateFlightResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	params := db.UpdateFlightParams{
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

	if _, err := queriesTx.UpdateFlight(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlight404Response{}, nil
		}
		return nil, err
	}

	flight, err := queriesTx.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlight404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateFlight200JSONResponse(fromDBFlight(flight)), nil
}

func (h *Handler) DeleteFlight(ctx context.Context, request api.DeleteFlightRequestObject) (api.DeleteFlightResponseObject, error) {
	if err := h.queries.DeleteFlight(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteFlight204Response{}, nil
}

func (h *Handler) DeleteAllFlights(ctx context.Context, request api.DeleteAllFlightsRequestObject) (api.DeleteAllFlightsResponseObject, error) {
	if err := h.queries.DeleteAllFlights(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllFlights204Response{}, nil
}
