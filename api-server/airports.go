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
	"github.com/stellora/airline/api-server/extdata"
)

func getAirportBySpec(ctx context.Context, queries *db.Queries, spec api.AirportSpec) (db.Airport, error) {
	if id, err := spec.AsAirportSpec0(); err == nil {
		return queries.GetAirport(ctx, int64(id))
	}
	if iataCode, err := spec.AsAirportSpec1(); err == nil {
		return queries.GetAirportByIATACode(ctx, iataCode)
	}
	panic("invalid AirportSpec")
}

func newAirportSpec(id int, iataCode string) api.AirportSpec {
	var spec api.AirportSpec
	if id != 0 {
		spec.FromAirportSpec0(id)
	} else {
		spec.FromAirportSpec1(iataCode)
	}
	return spec
}

// TODO!(sqs): remove? is unused
func airportSpecFromPathArg(arg string) api.AirportSpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return newAirportSpec(id, "")
	}
	return newAirportSpec(0, arg)
}

func fromDBAirport(a db.Airport) api.Airport {
	b := api.Airport{
		Id:       int(a.ID),
		IataCode: a.IataCode,
	}
	if a.OadbID.Valid {
		info := extdata.Airports.AirportByOAID(int(a.OadbID.Int64))
		b.Name = info.Airport.Name
		b.Country = string(info.Country.Code)
		b.Region = info.Region.Name
		b.Point = api.Point{
			Latitude:  info.Airport.LatitudeDeg,
			Longitude: info.Airport.LongitudeDeg,
		}
	}
	// TODO(sqs): handle case where there is no oadb_id or the OA database has no airport with the ID
	return b
}

func (h *Handler) GetAirport(ctx context.Context, request api.GetAirportRequestObject) (api.GetAirportResponseObject, error) {
	airport, err := h.queries.GetAirport(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetAirport404Response{}, nil
		}
	}
	return api.GetAirport200JSONResponse(fromDBAirport(airport)), nil
}

func (h *Handler) ListAirports(ctx context.Context, request api.ListAirportsRequestObject) (api.ListAirportsResponseObject, error) {
	airports, err := h.queries.ListAirports(ctx)
	if err != nil {
		return nil, err
	}
	return api.ListAirports200JSONResponse(mapSlice(fromDBAirport, airports)), nil
}

var validAirportIATACode = regexp.MustCompile(`^[A-Z]{3}$`)

func (h *Handler) CreateAirport(ctx context.Context, request api.CreateAirportRequestObject) (api.CreateAirportResponseObject, error) {
	if !validAirportIATACode.MatchString(request.Body.IataCode) {
		log.Println("invalid IATA") // TODO(sqs): return error
		return api.CreateAirport400Response{}, nil
	}

	params := db.CreateAirportParams{
		IataCode: request.Body.IataCode,
	}
	if info := extdata.Airports.AirportByIATACode(request.Body.IataCode); info != nil {
		params.OadbID = sql.NullInt64{Int64: int64(info.Airport.ID), Valid: true}
	}
	created, err := h.queries.CreateAirport(ctx, params)
	if err != nil {
		log.Println(err) // TODO(sqs): return error
		return api.CreateAirport400Response{}, nil
	}
	return api.CreateAirport201JSONResponse(fromDBAirport(created)), nil
}

func (h *Handler) UpdateAirport(ctx context.Context, request api.UpdateAirportRequestObject) (api.UpdateAirportResponseObject, error) {
	params := db.UpdateAirportParams{
		ID: int64(request.Id),
	}
	if request.Body.IataCode != nil {
		params.IataCode = sql.NullString{String: *request.Body.IataCode, Valid: true}
	}

	updated, err := h.queries.UpdateAirport(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAirport404Response{}, nil
		}
		return nil, err
	}
	return api.UpdateAirport200JSONResponse(fromDBAirport(updated)), nil
}

func (h *Handler) DeleteAirport(ctx context.Context, request api.DeleteAirportRequestObject) (api.DeleteAirportResponseObject, error) {
	if err := h.queries.DeleteAirport(ctx, int64(request.Id)); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}
	return api.DeleteAirport204Response{}, nil
}

func (h *Handler) DeleteAllAirports(ctx context.Context, request api.DeleteAllAirportsRequestObject) (api.DeleteAllAirportsResponseObject, error) {
	if err := h.queries.DeleteAllAirports(ctx); err != nil {
		return nil, err
	}
	return api.DeleteAllAirports204Response{}, nil
}
