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

func fromDBSchedule(a db.SchedulesView) api.Schedule {
	daysOfWeek, _ := parseDaysOfWeek(a.DaysOfWeek)
	airline := fromDBAirline(db.Airline{
		ID:       a.AirlineID,
		IataCode: a.AirlineIataCode,
		Name:     a.AirlineName,
	})
	fleet := fromDBFleet(db.FleetsView{
		ID:          a.FleetID,
		AirlineID:   a.FleetAirlineID,
		Code:        a.FleetCode,
		Description: a.FleetDescription,
	})
	fleet.Airline = airline
	b := api.Schedule{
		Id:      int(a.ID),
		Airline: airline,
		Number:  a.Number,
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
		Fleet:         fleet,
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

func (h *Handler) GetSchedule(ctx context.Context, request api.GetScheduleRequestObject) (api.GetScheduleResponseObject, error) {
	flight, err := h.queries.GetSchedule(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetSchedule404Response{}, nil
		}
	}
	return api.GetSchedule200JSONResponse(fromDBSchedule(flight)), nil
}

func (h *Handler) ListSchedules(ctx context.Context, request api.ListSchedulesRequestObject) (api.ListSchedulesResponseObject, error) {
	flights, err := h.queries.ListSchedules(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListSchedules200JSONResponse(mapSlice(fromDBSchedule, flights)), nil
}

func (h *Handler) CreateSchedule(ctx context.Context, request api.CreateScheduleRequestObject) (api.CreateScheduleResponseObject, error) {
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

	fleet, err := getFleetBySpec(ctx, queriesTx, airline.ID, request.Body.Fleet)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("fleet %q not found", request.Body.Fleet)
		}
		return nil, fmt.Errorf("looking up fleet: %w", err)
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

	created, err := queriesTx.CreateSchedule(ctx, db.CreateScheduleParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		FleetID:              fleet.ID,
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

	flight, err := queriesTx.GetSchedule(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := syncScheduleInstances(ctx, queriesTx, flight); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateSchedule201JSONResponse(fromDBSchedule(flight)), nil
}

func (h *Handler) UpdateSchedule(ctx context.Context, request api.UpdateScheduleRequestObject) (api.UpdateScheduleResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	existing, err := queriesTx.GetSchedule(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateSchedule404Response{}, nil
		}
		return nil, err
	}

	params := db.UpdateScheduleParams{
		ID: int64(request.Id),
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
	if request.Body.Fleet != nil {
		fleet, err := getFleetBySpec(ctx, queriesTx, existing.AirlineID, *request.Body.Fleet)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("fleet %q not found", request.Body.Fleet)
			}
			return nil, fmt.Errorf("looking up fleet: %w", err)
		}
		params.FleetID = sql.NullInt64{Int64: fleet.ID, Valid: true}
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

	if _, err := queriesTx.UpdateSchedule(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateSchedule404Response{}, nil
		}
		return nil, err
	}

	flight, err := queriesTx.GetSchedule(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateSchedule404Response{}, nil
		}
		return nil, err
	}

	if err := syncScheduleInstances(ctx, queriesTx, flight); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateSchedule200JSONResponse(fromDBSchedule(flight)), nil
}

func (h *Handler) DeleteSchedule(ctx context.Context, request api.DeleteScheduleRequestObject) (api.DeleteScheduleResponseObject, error) {
	if err := h.queries.DeleteSchedule(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteSchedule204Response{}, nil
}

func (h *Handler) DeleteAllSchedules(ctx context.Context, request api.DeleteAllSchedulesRequestObject) (api.DeleteAllSchedulesResponseObject, error) {
	if err := h.queries.DeleteAllSchedules(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllSchedules204Response{}, nil
}
