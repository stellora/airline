package main

import (
	"context"
	"database/sql"
	"errors"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBFlightInstance(a db.FlightInstance, as db.FlightSchedulesView, aa *db.AircraftView) api.FlightInstance {
	// TODO!(sqs): bring in aircraft
	b := api.FlightInstance{
		Id:           int(a.ID),
		Source:       fromDBFlightSchedule(as),
		InstanceDate: openapi_types.Date{Time: a.InstanceDate},
	}
	if aa != nil {
		b.Aircraft = ptrTo(fromDBAircraft(*aa))
	}
	if a.Notes != "" {
		b.Notes = &a.Notes
	}
	return b
}

func fromDBFlightInstances(as []db.ListFlightInstancesRow) []api.FlightInstance {
	bs := make([]api.FlightInstance, len(as))
	for i, a := range as {
		bs[i] = fromDBFlightInstance(a.FlightInstance, a.FlightSchedulesView, nil)
	}
	return bs
}

func (h *Handler) GetFlightInstance(ctx context.Context, request api.GetFlightInstanceRequestObject) (api.GetFlightInstanceResponseObject, error) {
	row, err := h.queries.GetFlightInstance(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetFlightInstance404Response{}, nil
		}
	}
	return api.GetFlightInstance200JSONResponse(fromDBFlightInstance(row.FlightInstance, row.FlightSchedulesView, nil)), nil
}

func (h *Handler) ListFlightInstances(ctx context.Context, request api.ListFlightInstancesRequestObject) (api.ListFlightInstancesResponseObject, error) {
	rows, err := h.queries.ListFlightInstances(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListFlightInstances200JSONResponse(fromDBFlightInstances(rows)), nil
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

	return api.UpdateFlightInstance200JSONResponse(fromDBFlightInstance(row.FlightInstance, row.FlightSchedulesView, nil)), nil
}

func (h *Handler) DeleteFlightInstance(ctx context.Context, request api.DeleteFlightInstanceRequestObject) (api.DeleteFlightInstanceResponseObject, error) {
	if err := h.queries.DeleteFlightInstance(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteFlightInstance204Response{}, nil
}
