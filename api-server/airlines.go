package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"strconv"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func getAirlineBySpec(ctx context.Context, queries *db.Queries, spec api.AirlineSpec) (db.Airline, error) {
	if id, err := spec.AsAirlineSpec0(); err == nil {
		return queries.GetAirline(ctx, int64(id))
	}
	if iataCode, err := spec.AsAirlineIATACode(); err == nil {
		return queries.GetAirlineByIATACode(ctx, iataCode)
	}
	panic("invalid AirlineSpec")
}

func newAirlineSpec(id int, iataCode string) api.AirlineSpec {
	var spec api.AirlineSpec
	if id != 0 {
		spec.FromAirlineSpec0(id)
	} else {
		spec.FromAirlineIATACode(iataCode)
	}
	return spec
}

// TODO!(sqs): remove? is unused
func airlineSpecFromPathArg(arg string) api.AirlineSpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return newAirlineSpec(id, "")
	}
	return newAirlineSpec(0, arg)
}

func fromDBAirline(a db.Airline) api.Airline {
	b := api.Airline{
		Id:       int(a.ID),
		IataCode: a.IataCode,
		Name:     a.Name,
	}
	return b
}

func (h *Handler) GetAirline(ctx context.Context, request api.GetAirlineRequestObject) (api.GetAirlineResponseObject, error) {
	airline, err := h.queries.GetAirline(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetAirline404Response{}, nil
		}
	}
	return api.GetAirline200JSONResponse(fromDBAirline(airline)), nil
}

func (h *Handler) ListAirlines(ctx context.Context, request api.ListAirlinesRequestObject) (api.ListAirlinesResponseObject, error) {
	airlines, err := h.queries.ListAirlines(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListAirlines200JSONResponse(mapSlice(fromDBAirline, airlines)), nil
}

var validAirlineIATACode = regexp.MustCompile(`^[A-Z]{2}$`)

func (h *Handler) CreateAirline(ctx context.Context, request api.CreateAirlineRequestObject) (api.CreateAirlineResponseObject, error) {
	if !validAirlineIATACode.MatchString(request.Body.IataCode) {
		log.Println("invalid IATA") // TODO(sqs): return error
		return api.CreateAirline400Response{}, nil
	}

	params := db.CreateAirlineParams{
		IataCode: request.Body.IataCode,
		Name:     request.Body.Name,
	}
	created, err := h.queries.CreateAirline(ctx, params)
	if err != nil {
		log.Println(err) // TODO(sqs): return error
		return api.CreateAirline400Response{}, nil
	}
	return api.CreateAirline201JSONResponse(fromDBAirline(created)), nil
}

func (h *Handler) UpdateAirline(ctx context.Context, request api.UpdateAirlineRequestObject) (api.UpdateAirlineResponseObject, error) {
	params := db.UpdateAirlineParams{
		ID: int64(request.Id),
	}
	if request.Body.IataCode != nil {
		params.IataCode = sql.NullString{String: *request.Body.IataCode, Valid: true}
	}
	if request.Body.Name != nil {
		params.Name = sql.NullString{String: *request.Body.Name, Valid: true}
	}

	updated, err := h.queries.UpdateAirline(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAirline404Response{}, nil
		}
		return nil, err
	}
	return api.UpdateAirline200JSONResponse(fromDBAirline(updated)), nil
}

func (h *Handler) DeleteAirline(ctx context.Context, request api.DeleteAirlineRequestObject) (api.DeleteAirlineResponseObject, error) {
	if err := h.queries.DeleteAirline(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteAirline204Response{}, nil
}

func (h *Handler) DeleteAllAirlines(ctx context.Context, request api.DeleteAllAirlinesRequestObject) (api.DeleteAllAirlinesResponseObject, error) {
	if err := h.queries.DeleteAllAirlines(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllAirlines204Response{}, nil
}
