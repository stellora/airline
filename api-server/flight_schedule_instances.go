package main

import (
	"context"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func (h *Handler) ListFlightInstancesForFlightSchedule(ctx context.Context, request api.ListFlightInstancesForFlightScheduleRequestObject) (api.ListFlightInstancesForFlightScheduleResponseObject, error) {
	rows, err := h.queries.ListFlightInstancesForFlightSchedule(ctx, int64(request.Id))
	if err != nil {
		return nil, err
	}

	rows2 := make([]db.ListFlightInstancesRow, len(rows))
	for i, r := range rows {
		rows2[i] = db.ListFlightInstancesRow{FlightInstance: r.FlightInstance, FlightSchedulesView: r.FlightSchedulesView}
	}
	return api.ListFlightInstancesForFlightSchedule200JSONResponse(fromDBFlightInstances(rows2)), nil
}
