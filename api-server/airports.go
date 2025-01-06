package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
)

func getAirportBySpec(ctx context.Context, queries *db.Queries, spec api.AirportSpec) (db.Airport, error) {
	if id, err := spec.AsAirportID(); err == nil {
		return queries.GetAirport(ctx, int64(id))
	}
	if iataCode, err := spec.AsAirportIATACode(); err == nil {
		return queries.GetAirportByIATACode(ctx, iataCode)
	}
	panic("invalid AirportSpec")
}

// getOrCreateAirportBySpec is like getAirportBySpec, but it silently creates the airport if it does
// not exist and the spec is an airport IATA code.
func getOrCreateAirportBySpec(ctx context.Context, dbtx db.DBTX, queriesTx *db.Queries, spec api.AirportSpec) (db.Airport, error) {
	if id, err := spec.AsAirportID(); err == nil {
		return queriesTx.GetAirport(ctx, int64(id))
	}
	if iataCode, err := spec.AsAirportIATACode(); err == nil {
		// Ensure we're in a DB transaction.
		var (
			tx            *sql.Tx
			txCreatedByUs = false
		)
		if dbh, ok := dbtx.(*sql.DB); ok {
			tx, err = dbh.BeginTx(ctx, nil)
			if err != nil {
				return db.Airport{}, err
			}
			defer tx.Rollback()
			queriesTx = queriesTx.WithTx(tx)
			txCreatedByUs = true
		} else {
			// Already in a transaction.
			tx = dbtx.(*sql.Tx)
		}

		airport, err := queriesTx.GetAirportByIATACode(ctx, iataCode)
		if errors.Is(err, sql.ErrNoRows) {
			if extdata.Airports.AirportByIATACode(iataCode) == nil {
				return db.Airport{}, err
			}
			airport, err = createAirport(ctx, queriesTx, api.CreateAirportRequestObject{
				Body: &api.CreateAirportJSONRequestBody{
					IataCode: iataCode,
				},
			})
			if err != nil {
				return db.Airport{}, err
			}
		} else if err != nil {
			return db.Airport{}, err
		}

		// Only commit if we created the transaction.
		if txCreatedByUs {
			if err := tx.Commit(); err != nil {
				return db.Airport{}, err
			}
		}

		return airport, nil
	}
	panic("invalid AirportSpec")
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
		b.TimezoneID = info.Airport.TimezoneID
	}
	// TODO(sqs): handle case where there is no oadb_id or the OA database has no airport with the ID
	return b
}

func (h *Handler) GetAirport(ctx context.Context, request api.GetAirportRequestObject) (api.GetAirportResponseObject, error) {
	airport, err := getAirportBySpec(ctx, h.queries, request.AirportSpec)
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
	created, err := createAirport(ctx, h.queries, request)
	if err != nil {
		return nil, err
	}
	return api.CreateAirport201JSONResponse(fromDBAirport(created)), nil
}

func createAirport(ctx context.Context, queriesTx *db.Queries, request api.CreateAirportRequestObject) (db.Airport, error) {
	if !validAirportIATACode.MatchString(request.Body.IataCode) {
		return db.Airport{}, fmt.Errorf("invalid IATA code: %s", request.Body.IataCode)
	}

	params := db.CreateAirportParams{
		IataCode: request.Body.IataCode,
	}
	if info := extdata.Airports.AirportByIATACode(request.Body.IataCode); info != nil {
		params.OadbID = sql.NullInt64{Int64: int64(info.Airport.ID), Valid: true}
	}
	created, err := queriesTx.CreateAirport(ctx, params)
	if err != nil {
		return db.Airport{}, fmt.Errorf("failed to create airport %q: %w", request.Body.IataCode, err)
	}
	return created, nil
}

func (h *Handler) UpdateAirport(ctx context.Context, request api.UpdateAirportRequestObject) (api.UpdateAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airport, err := getAirportBySpec(ctx, queriesTx, request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAirport404Response{}, nil
		}
		return nil, err
	}

	params := db.UpdateAirportParams{ID: airport.ID}
	if request.Body.IataCode != nil {
		params.IataCode = sql.NullString{String: *request.Body.IataCode, Valid: true}
		info := extdata.Airports.AirportByIATACode(*request.Body.IataCode)
		if info != nil {
			params.OadbID = sql.NullInt64{Int64: int64(info.Airport.ID), Valid: true}
		} else {
			params.OadbID = sql.NullInt64{Valid: false}
		}
	}

	updated, err := queriesTx.UpdateAirport(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.UpdateAirport404Response{}, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.UpdateAirport200JSONResponse(fromDBAirport(updated)), nil
}

func (h *Handler) DeleteAirport(ctx context.Context, request api.DeleteAirportRequestObject) (api.DeleteAirportResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	airport, err := getAirportBySpec(ctx, queriesTx, request.AirportSpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeleteAirport404Response{}, nil
		}
		return nil, err
	}

	if err := queriesTx.DeleteAirport(ctx, airport.ID); err != nil {
		// TODO(sqs): check if it was actually deleted
		return nil, err
	}

	if err := tx.Commit(); err != nil {
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
