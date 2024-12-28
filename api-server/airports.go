package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"regexp"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
)

func getAirport(id int) *api.Airport {
	panic("TODO!(sqs)")
}

func getAirportBySpec(spec api.AirportSpec) *api.Airport {
	if id, err := spec.AsAirportSpec0(); err == nil {
		return getAirport(id)
	}
	if iataCode, err := spec.AsAirportSpec1(); err == nil {
		for _, airport := range airports {
			if airport.IataCode == iataCode {
				return airport
			}
		}
	}
	return nil
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

var validIATACode = regexp.MustCompile(`^[A-Z]{3}$`)

func (h *Handler) CreateAirport(ctx context.Context, request api.CreateAirportRequestObject) (api.CreateAirportResponseObject, error) {
	if !validIATACode.MatchString(request.Body.IataCode) {
		log.Println("invalid IATA") // TODO(sqs): return error
		return api.CreateAirport400Response{}, nil
	}

	params := db.CreateAirportParams{
		IataCode: request.Body.IataCode,
	}
	if info := extdata.Airports.AirportByIATACode(request.Body.IataCode); info != nil {
		params.OadbID = sql.NullInt64{Int64: int64(info.Airport.ID), Valid: true}
	}
	if _, err := h.queries.CreateAirport(ctx, params); err != nil {
		log.Println(err) // TODO(sqs): return error
		return api.CreateAirport400Response{}, nil
	}
	return api.CreateAirport201Response{}, nil
}

func (h *Handler) UpdateAirport(ctx context.Context, request api.UpdateAirportRequestObject) (api.UpdateAirportResponseObject, error) {
	airport := getAirport(request.Id)
	if airport == nil {
		return &api.UpdateAirport404Response{}, nil
	}

	if request.Body.IataCode != nil {
		airport.IataCode = *request.Body.IataCode
	}
	return api.UpdateAirport204Response{}, nil
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
