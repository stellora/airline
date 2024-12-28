package main

import (
	"context"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBRoute(a db.ListRoutesRow) api.Route {
	return api.Route{
		OriginAirport:      fromDBAirport(a.OriginAirport),
		DestinationAirport: fromDBAirport(a.DestinationAirport),
		FlightsCount:       int(a.FlightsCount),
	}
}

func (h *Handler) ListRoutes(ctx context.Context, request api.ListRoutesRequestObject) (api.ListRoutesResponseObject, error) {
	routes, err := h.queries.ListRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListRoutes200JSONResponse(mapSlice(fromDBRoute, routes)), nil
}
