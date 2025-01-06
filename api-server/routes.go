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
		SchedulesCount: int(a.SchedulesCount),
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

func getRouteAirports(ctx context.Context, tx db.DBTX, queriesTx *db.Queries, route string) (origin, destination db.Airport, err error) {
	originIATACode, destinationIATACode, err := parseRoute(route)
	if err != nil {
		return
	}
	origin, err = getOrCreateAirportBySpec(ctx, tx, queriesTx, api.NewAirportSpec(0, originIATACode))
	if err != nil {
		return
	}
	destination, err = getOrCreateAirportBySpec(ctx, tx, queriesTx, api.NewAirportSpec(0, destinationIATACode))
	if err != nil {
		return
	}
	return origin, destination, nil
}

func (h *Handler) ListSchedulesByRoute(ctx context.Context, request api.ListSchedulesByRouteRequestObject) (api.ListSchedulesByRouteResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	origin, destination, err := getRouteAirports(ctx, tx, queriesTx, request.Route)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListSchedulesByRoute404Response{}, nil
		}
		return nil, err
	}

	schedules, err := queriesTx.ListSchedulesByRoute(ctx, db.ListSchedulesByRouteParams{
		OriginAirport:      origin.ID,
		DestinationAirport: destination.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListSchedulesByRoute200JSONResponse(mapSlice(fromDBSchedule, schedules)), nil
}

func (h *Handler) ListFlightsByRoute(ctx context.Context, request api.ListFlightsByRouteRequestObject) (api.ListFlightsByRouteResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	origin, destination, err := getRouteAirports(ctx, tx, queriesTx, request.Route)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListFlightsByRoute404Response{}, nil
		}
		return nil, err
	}

	flights, err := queriesTx.ListFlightsByRoute(ctx, db.ListFlightsByRouteParams{
		OriginAirport:      origin.ID,
		DestinationAirport: destination.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListFlightsByRoute200JSONResponse(mapSlice(fromDBFlight, flights)), nil
}
