package main

import (
	"context"
	"database/sql"
	"errors"
	"regexp"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBRoute(a db.Route) api.Route {
	b := api.Route{
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
		FlightSchedulesCount: int(a.FlightSchedulesCount),
	}
	b.DistanceMiles = distanceMilesBetweenAirports(b.OriginAirport, b.DestinationAirport)
	return b
}

var (
	validRoute      = regexp.MustCompile(`^([A-Z]{3})-([A-Z]{3})$`)
	errInvalidRoute = errors.New("invalid route")
)

// parseRoute parses a route string of the form "AAA-BBB".
func parseRoute(route string) (origin, destination string, err error) {
	matches := validRoute.FindStringSubmatch(route)
	if matches == nil {
		return "", "", errInvalidRoute
	}
	origin, destination = matches[1], matches[2]
	return
}

func (h *Handler) ListRoutes(ctx context.Context, request api.ListRoutesRequestObject) (api.ListRoutesResponseObject, error) {
	routes, err := h.queries.ListRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListRoutes200JSONResponse(mapSlice(fromDBRoute, routes)), nil
}

func (h *Handler) GetRoute(ctx context.Context, request api.GetRouteRequestObject) (api.GetRouteResponseObject, error) {
	originIATACode, destinationIATACode, err := parseRoute(request.Route)
	if err != nil {
		return nil, err
	}

	route, err := h.queries.GetRouteByIATACodes(ctx, db.GetRouteByIATACodesParams{
		OriginAirportIataCode:      originIATACode,
		DestinationAirportIataCode: destinationIATACode,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetRoute404Response{}, nil
		}
		return nil, err
	}

	return api.GetRoute200JSONResponse(fromDBRoute(route)), nil
}
