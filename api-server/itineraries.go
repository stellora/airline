package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func getItineraryBySpec(ctx context.Context, queriesTx *db.Queries, spec api.ItinerarySpec) (db.Itinerary, error) {
	if id, err := spec.AsItineraryID(); err == nil {
		return queriesTx.GetItinerary(ctx, int64(id))
	}
	if recordLocator, err := spec.AsRecordLocator(); err == nil {
		return queriesTx.GetItineraryByRecordLocator(ctx, recordLocator)
	}
	panic("invalid ItinerarySpec: " + fmt.Sprintf("%#v", spec))
}

func fromDBItinerary(ctx context.Context, queriesTx *db.Queries, i db.Itinerary) (api.Itinerary, error) {
	flights, err := queriesTx.ListItineraryFlights(ctx, i.ID)
	if err != nil {
		return api.Itinerary{}, err
	}

	passengers, err := queriesTx.ListItineraryPassengers(ctx, i.ID)
	if err != nil {
		return api.Itinerary{}, err
	}

	return api.Itinerary{
		Id:         int(i.ID),
		RecordID:   i.RecordID,
		Flights:    mapSlice(fromDBFlightInstance, flights),
		Passengers: mapSlice(fromDBPassenger, passengers),
	}, nil
}

func (h *Handler) GetItinerary(ctx context.Context, request api.GetItineraryRequestObject) (api.GetItineraryResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	itinerary, err := getItineraryBySpec(ctx, queriesTx, request.ItinerarySpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetItinerary404Response{}, nil
		}
		return nil, err
	}

	result, err := fromDBItinerary(ctx, queriesTx, itinerary)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.GetItinerary200JSONResponse(result), nil
}

func (h *Handler) ListItineraries(ctx context.Context, request api.ListItinerariesRequestObject) (api.ListItinerariesResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	itineraries, err := queriesTx.ListItineraries(ctx)
	if err != nil {
		return nil, err
	}

	var result []api.Itinerary
	for _, i := range itineraries {
		converted, err := fromDBItinerary(ctx, queriesTx, i)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.ListItineraries200JSONResponse(result), nil
}

func (h *Handler) CreateItinerary(ctx context.Context, request api.CreateItineraryRequestObject) (api.CreateItineraryResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	// Generate a unique record locator
	recordLocator := generateRecordLocator()
	created, err := queriesTx.CreateItinerary(ctx, recordLocator)
	if err != nil {
		return nil, err
	}

	// Add flights
	for _, flightID := range request.Body.FlightInstanceIDs {
		err := queriesTx.AddFlightToItinerary(ctx, db.AddFlightToItineraryParams{
			ItineraryID:      created.ID,
			FlightInstanceID: int64(flightID),
		})
		if err != nil {
			return nil, err
		}
	}

	// Add passengers
	for _, passengerID := range request.Body.PassengerIDs {
		err := queriesTx.AddPassengerToItinerary(ctx, db.AddPassengerToItineraryParams{
			ItineraryID: created.ID,
			PassengerID: int64(passengerID),
		})
		if err != nil {
			return nil, err
		}
	}

	result, err := fromDBItinerary(ctx, queriesTx, created)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.CreateItinerary201JSONResponse(result), nil
}

func (h *Handler) DeleteItinerary(ctx context.Context, request api.DeleteItineraryRequestObject) (api.DeleteItineraryResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	itinerary, err := getItineraryBySpec(ctx, queriesTx, request.ItinerarySpec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.DeleteItinerary404Response{}, nil
		}
		return nil, err
	}

	if err := queriesTx.DeleteItinerary(ctx, itinerary.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return api.DeleteItinerary204Response{}, nil
}

var testingDummyRecordLocators []string

// generateRecordLocator generates a unique 6-character string matching the `[A-Z0-9]{6}“ pattern.
func generateRecordLocator() string {
	if testingDummyRecordLocators != nil {
		v := fmt.Sprintf("TEST%02d", len(testingDummyRecordLocators))
		testingDummyRecordLocators = append(testingDummyRecordLocators, v)
		return v
	}

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
