package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBSeatAssignment(a db.SeatAssignmentsView) api.SeatAssignment {
	return api.SeatAssignment{
		Id:        int(a.ID),
		SegmentID: int(a.SegmentID),
		Itinerary: api.ItinerarySpecs{
			Id:       int(a.ItineraryID),
			RecordID: a.ItineraryRecordID,
		},
		Passenger: api.Passenger{
			Id:   int(a.PassengerID),
			Name: a.PassengerName,
		},
		FlightID: int(a.FlightID),
		Seat:     a.Seat,
	}
}

func (h *Handler) ListSeatAssignmentsForFlight(ctx context.Context, request api.ListSeatAssignmentsForFlightRequestObject) (api.ListSeatAssignmentsForFlightResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	if _, err := queriesTx.GetFlight(ctx, int64(request.FlightID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListSeatAssignmentsForFlight404Response{}, nil
		}
		return nil, err
	}

	assignments, err := queriesTx.ListSeatAssignmentsForFlight(ctx, int64(request.FlightID))
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListSeatAssignmentsForFlight200JSONResponse(mapSlice(fromDBSeatAssignment, assignments)), nil
}

func (h *Handler) SetSeatAssignment(ctx context.Context, request api.SetSeatAssignmentRequestObject) (api.SetSeatAssignmentResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	if _, err := getItineraryBySpec(ctx, queriesTx, request.ItinerarySpec); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.SetSeatAssignment404Response{}, nil
		}
		return nil, err

	}

	segment, err := getSegmentBySpec(ctx, queriesTx, request.SegmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.SetSeatAssignment404Response{}, nil
		}
		return nil, err
	}

	params := db.CreateSeatAssignmentParams{
		SegmentID:   int64(request.SegmentID),
		FlightID:    int64(segment.FlightID),
		PassengerID: int64(request.PassengerID),
		Seat:        request.Body.Seat,
	}

	created, err := queriesTx.CreateSeatAssignment(ctx, params)
	if err != nil {
		return nil, err
	}

	// TODO!(sqs): handle setting new seat assignment when one exists

	assignment, err := queriesTx.GetSeatAssignment(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.SetSeatAssignment201JSONResponse(fromDBSeatAssignment(assignment)), nil
}
