package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
	"github.com/stellora/airline/api-server/zonedtime"
)

func fromDBFlight(a db.FlightsView) api.Flight {
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
	b := api.Flight{
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
		Fleet:             fleet,
		DepartureDateTime: a.DepartureDatetime,
		ArrivalDateTime:   a.ArrivalDatetime,
		Notes:             a.Notes,
		Published:         a.Published,
	}
	b.DistanceMiles = distanceMilesBetweenAirports(b.OriginAirport, b.DestinationAirport)
	if a.AircraftID.Valid {
		b.Aircraft = ptrTo(fromDBAircraft(db.AircraftView{
			ID:              a.AircraftID.Int64,
			Registration:    a.AircraftRegistration.String,
			AircraftType:    a.AircraftAircraftType.String,
			AirlineID:       a.AircraftAirlineID.Int64,
			AirlineIataCode: a.AircraftAirlineIataCode.String,
			AirlineName:     a.AircraftAirlineName.String,
		}))
	}

	if a.SourceScheduleID.Valid {
		b.ScheduleID = ptrTo(int(a.SourceScheduleID.Int64))
		b.ScheduleInstanceDate = ptrTo(a.SourceScheduleInstanceLocaldate.String())
	}

	return b
}

func (h *Handler) GetFlight(ctx context.Context, request api.GetFlightRequestObject) (api.GetFlightResponseObject, error) {
	row, err := h.queries.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.GetFlight404Response{}, nil
		}
	}
	return api.GetFlight200JSONResponse(fromDBFlight(row)), nil
}

func (h *Handler) ListFlights(ctx context.Context, request api.ListFlightsRequestObject) (api.ListFlightsResponseObject, error) {
	rows, err := h.queries.ListFlights(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListFlights200JSONResponse(mapSlice(fromDBFlight, rows)), nil
}

// Ensure departure/arrival datetimes use the locations of the departure/arrival airports,
// respectively.
func checkDepartureArrivalDateTimesMatchAirportTimezones(departure, arrival *zonedtime.ZonedTime, origin, destination db.Airport) error {
	dateTimeMatchesAirport := func(t zonedtime.ZonedTime, a db.Airport) (wantTzID string, matches bool) {
		info := extdata.Airports.AirportByOAID(int(a.OadbID.Int64))
		if t.Location().String() != info.Airport.TimezoneID {
			return info.Airport.TimezoneID, false
		}
		return "", true
	}

	if departure != nil {
		if wantTz, ok := dateTimeMatchesAirport(*departure, origin); !ok {
			return fmt.Errorf("departureDateTime must match timezone of origin airport %s: %q != %q", origin.IataCode, departure.Location(), wantTz)
		}
	}
	if arrival != nil {
		if wantTz, ok := dateTimeMatchesAirport(*arrival, destination); !ok {
			return fmt.Errorf("arrivalDateTime must match timezone of destination airport %s: %q != %q", destination.IataCode, arrival.Location(), wantTz)
		}
	}
	return nil
}

func (h *Handler) CreateFlight(ctx context.Context, request api.CreateFlightRequestObject) (api.CreateFlightResponseObject, error) {
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

	var aircraftID sql.NullInt64
	if request.Body.Aircraft != nil {
		aircraft, err := getAircraftBySpec(ctx, queriesTx, *request.Body.Aircraft)
		if err != nil {
			return nil, err
		}
		aircraftID = sql.NullInt64{Valid: true, Int64: aircraft.ID}
	}

	if err := checkDepartureArrivalDateTimesMatchAirportTimezones(&request.Body.DepartureDateTime, &request.Body.ArrivalDateTime, originAirport, destinationAirport); err != nil {
		return nil, err
	}

	created, err := queriesTx.CreateFlight(ctx, db.CreateFlightParams{
		AirlineID:            airline.ID,
		Number:               request.Body.Number,
		OriginAirportID:      originAirport.ID,
		DestinationAirportID: destinationAirport.ID,
		FleetID:              fleet.ID,
		AircraftID:           aircraftID,
		DepartureDatetime:    &request.Body.DepartureDateTime,
		ArrivalDatetime:      &request.Body.ArrivalDateTime,
		DepartureDatetimeUtc: request.Body.DepartureDateTime.Time.In(time.UTC),
		ArrivalDatetimeUtc:   request.Body.ArrivalDateTime.Time.In(time.UTC),
		Notes:                request.Body.Notes,
		Published:            request.Body.Published != nil && *request.Body.Published,
	})
	if err != nil {
		return nil, err
	}

	row, err := queriesTx.GetFlight(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateFlight201JSONResponse(fromDBFlight(row)), nil
}

func (h *Handler) UpdateFlight(ctx context.Context, request api.UpdateFlightRequestObject) (api.UpdateFlightResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	existing, err := queriesTx.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlight404Response{}, nil
		}
		return nil, err
	}

	params := db.UpdateFlightParams{
		ID: int64(request.Id),
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

	// TODO!(sqs): update other fields

	// TODO!(sqs): when doing this, ensure we use either the old or new airports, if theyre being updated in this call
	// if err := checkDepartureArrivalDateTimesMatchAirportTimezones(request.Body.DepartureDateTime, request.Body.ArrivalDateTime, originAirport, destinationAirport); err != nil {
	// 	return nil, err
	// }

	if _, err := queriesTx.UpdateFlight(ctx, params); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlight404Response{}, nil
		}
		return nil, err
	}

	row, err := queriesTx.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateFlight404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateFlight200JSONResponse(fromDBFlight(row)), nil
}

func (h *Handler) DeleteFlight(ctx context.Context, request api.DeleteFlightRequestObject) (api.DeleteFlightResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	row, err := queriesTx.GetFlight(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.DeleteFlight404Response{}, nil
		}
		return nil, err
	}

	// Only flights created from manual input can be deleted. To delete a flights
	// created from a schedule, you need to edit the schedule so that it deletes the
	// instance when it syncs the new schedule definition.
	if row.SourceScheduleID.Valid {
		return api.DeleteFlight400Response{}, nil
	}

	if err := queriesTx.DeleteFlight(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.DeleteFlight204Response{}, nil
}
