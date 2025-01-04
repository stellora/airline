package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/localtime"
)

func fromDBFlightSchedule(a db.FlightSchedulesView) api.FlightSchedule {
	daysOfWeek, _ := parseDaysOfWeek(a.DaysOfWeek)
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
		AircraftType:  fromAircraftTypeCode(a.AircraftType),
		StartDate:     a.StartLocaldate.String(),
		EndDate:       a.EndLocaldate.String(),
		DaysOfWeek:    daysOfWeek,
		DepartureTime: a.DepartureLocaltime.String(),
		DurationSec:   int(a.DurationSec),
		Published:     a.Published,
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

	if request.Body.AircraftType == "" {
		return nil, errors.New("aircraftType must not be empty")
	}

	startDate, err := localtime.ParseLocalDate(request.Body.StartDate)
	if err != nil {
		return nil, fmt.Errorf("parsing startDate: %w", err)
	}
	endDate, err := localtime.ParseLocalDate(request.Body.EndDate)
	if err != nil {
		return nil, fmt.Errorf("parsing endDate: %w", err)
	}

	departureTime, err := localtime.ParseTimeOfDay(request.Body.DepartureTime)
	if err != nil {
		return nil, fmt.Errorf("parsing departureTime: %w", err)
	}

	created, err := queriesTx.CreateFlightSchedule(ctx, db.CreateFlightScheduleParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		AircraftType:         request.Body.AircraftType,
		StartLocaldate:       &startDate,
		EndLocaldate:         &endDate,
		DaysOfWeek:           toDBDaysOfWeek(request.Body.DaysOfWeek),
		DepartureLocaltime:   &departureTime,
		DurationSec:          int64(request.Body.DurationSec),
		Published:            request.Body.Published != nil && *request.Body.Published,
	})
	if err != nil {
		return nil, err
	}

	flight, err := queriesTx.GetFlightSchedule(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := syncFlightScheduleInstances(ctx, queriesTx, flight); err != nil {
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
	if request.Body.AircraftType != nil {
		params.AircraftType = sql.NullString{String: *request.Body.AircraftType, Valid: true}
	}
	if request.Body.StartDate != nil {
		startDate, err := localtime.ParseLocalDate(*request.Body.StartDate)
		if err != nil {
			return nil, fmt.Errorf("parsing startDate: %w", err)
		}
		params.StartLocaldate = &startDate
	}
	if request.Body.EndDate != nil {
		endDate, err := localtime.ParseLocalDate(*request.Body.EndDate)
		if err != nil {
			return nil, fmt.Errorf("parsing endDate: %w", err)
		}
		params.EndLocaldate = &endDate
	}
	if request.Body.DaysOfWeek != nil {
		params.DaysOfWeek = sql.NullString{String: toDBDaysOfWeek(*request.Body.DaysOfWeek), Valid: true}
	}
	if request.Body.DepartureTime != nil {
		departureTime, err := localtime.ParseTimeOfDay(*request.Body.DepartureTime)
		if err != nil {
			return nil, fmt.Errorf("parsing departureTime: %w", err)
		}
		params.DepartureLocaltime = &departureTime
	}
	if request.Body.DurationSec != nil {
		params.DurationSec = sql.NullInt64{Int64: int64(*request.Body.DurationSec), Valid: true}
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

	if err := syncFlightScheduleInstances(ctx, queriesTx, flight); err != nil {
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
