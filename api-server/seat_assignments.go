package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBSeatAssignment(s db.SeatAssignment) api.SeatAssignment {
	return api.SeatAssignment{
		Id:               int(s.ID),
		ItineraryID:      int(s.ItineraryID),
		PassengerID:      int(s.PassengerID),
		FlightInstanceID: int(s.FlightInstanceID),
		Seat:             s.Seat,
	}
}

func (h *Handler) ListSeatAssignmentsForFlightInstance(ctx context.Context, request api.ListSeatAssignmentsForFlightInstanceRequestObject) (api.ListSeatAssignmentsForFlightInstanceResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	if _, err := queriesTx.GetFlightInstance(ctx, int64(request.FlightInstanceID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListSeatAssignmentsForFlightInstance404Response{}, nil
		}
		return nil, err
	}

	assignments, err := queriesTx.ListSeatAssignmentsForFlightInstance(ctx, int64(request.FlightInstanceID))
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListSeatAssignmentsForFlightInstance200JSONResponse(mapSlice(fromDBSeatAssignment, assignments)), nil
}

func (h *Handler) CreateSeatAssignment(ctx context.Context, request api.CreateSeatAssignmentRequestObject) (api.CreateSeatAssignmentResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	params := db.CreateSeatAssignmentParams{
		ItineraryID:      int64(request.Body.ItineraryID),
		PassengerID:      int64(request.Body.PassengerID),
		FlightInstanceID: int64(request.FlightInstanceID),
		Seat:             request.Body.Seat,
	}

	created, err := queriesTx.CreateSeatAssignment(ctx, params)
	if err != nil {
		return nil, err
	}

	assignment, err := queriesTx.GetSeatAssignment(ctx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateSeatAssignment201JSONResponse(fromDBSeatAssignment(assignment)), nil
}
