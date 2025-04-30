package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func fromDBAirlineLoyalty(a db.AirlineLoyalty) api.AirlineLoyalty {
	return api.AirlineLoyalty{
		Id:             int(a.ID),
		Airline:        fromDBAirline(a.Airline),
		PassengerId:    int(a.PassengerID),
		MileageBalance: int(a.MileageBalance),
	}
}

// GetAirlineLoyaltiesByPassenger retrieves all airline loyalty programs for a specific passenger
func (h *Handler) GetAirlineLoyaltiesByPassenger(ctx context.Context, passengerID int64) ([]api.AirlineLoyalty, error) {
	loyalties, err := h.queries.GetAirlineLoyaltiesByPassenger(ctx, passengerID)
	if err != nil {
		return nil, err
	}

	result := make([]api.AirlineLoyalty, len(loyalties))
	for i, loyalty := range loyalties {
		airline, err := h.queries.GetAirline(ctx, loyalty.AirlineID)
		if err != nil {
			return nil, err
		}
		loyalty.Airline = airline
		result[i] = fromDBAirlineLoyalty(loyalty)
	}

	return result, nil
}

// CreateOrUpdateAirlineLoyalty creates or updates an airline loyalty program entry
func (h *Handler) CreateOrUpdateAirlineLoyalty(ctx context.Context, airlineID int64, passengerID int64, mileageToAdd int) (api.AirlineLoyalty, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return api.AirlineLoyalty{}, err
	}
	defer tx.Rollback()
	queriesTx := h.queries.WithTx(tx)

	// Try to get existing loyalty program entry
	existing, err := queriesTx.GetAirlineLoyalty(ctx, db.GetAirlineLoyaltyParams{
		AirlineID:   airlineID,
		PassengerID: passengerID,
	})

	var id int64
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Create new loyalty program entry
			id, err = queriesTx.CreateAirlineLoyalty(ctx, db.CreateAirlineLoyaltyParams{
				AirlineID:      airlineID,
				PassengerID:    passengerID,
				MileageBalance: int64(mileageToAdd),
			})
			if err != nil {
				return api.AirlineLoyalty{}, err
			}
		} else {
			return api.AirlineLoyalty{}, err
		}
	} else {
		// Update existing loyalty program entry
		id, err = queriesTx.UpdateAirlineLoyalty(ctx, db.UpdateAirlineLoyaltyParams{
			ID:            existing.ID,
			MileageBalance: sql.NullInt64{Valid: true, Int64: existing.MileageBalance + int64(mileageToAdd)},
		})
		if err != nil {
			return api.AirlineLoyalty{}, err
		}
	}

	// Get updated loyalty program
	updated, err := queriesTx.GetAirlineLoyaltyByID(ctx, id)
	if err != nil {
		return api.AirlineLoyalty{}, err
	}

	airline, err := queriesTx.GetAirline(ctx, updated.AirlineID)
	if err != nil {
		return api.AirlineLoyalty{}, err
	}
	updated.Airline = airline

	if err := tx.Commit(); err != nil {
		return api.AirlineLoyalty{}, err
	}

	return fromDBAirlineLoyalty(updated), nil
}

// GetAirlineLoyalty gets a specific loyalty program entry
func (h *Handler) GetAirlineLoyalty(ctx context.Context, request api.GetAirlineLoyaltyRequestObject) (api.GetAirlineLoyaltyResponseObject, error) {
	loyalty, err := h.queries.GetAirlineLoyaltyByID(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetAirlineLoyalty404Response{}, nil
		}
		return nil, err
	}

	airline, err := h.queries.GetAirline(ctx, loyalty.AirlineID)
	if err != nil {
		return nil, err
	}
	loyalty.Airline = airline

	return api.GetAirlineLoyalty200JSONResponse(fromDBAirlineLoyalty(loyalty)), nil
}

// ListAirlineLoyalties lists all loyalty program entries
func (h *Handler) ListAirlineLoyalties(ctx context.Context, request api.ListAirlineLoyaltiesRequestObject) (api.ListAirlineLoyaltiesResponseObject, error) {
	loyalties, err := h.queries.ListAirlineLoyalties(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]api.AirlineLoyalty, len(loyalties))
	for i, loyalty := range loyalties {
		airline, err := h.queries.GetAirline(ctx, loyalty.AirlineID)
		if err != nil {
			return nil, err
		}
		loyalty.Airline = airline
		result[i] = fromDBAirlineLoyalty(loyalty)
	}

	return api.ListAirlineLoyalties200JSONResponse(result), nil
}

// ListPassengerAirlineLoyalties lists all loyalty program entries for a specific passenger
func (h *Handler) ListPassengerAirlineLoyalties(ctx context.Context, request api.ListPassengerAirlineLoyaltiesRequestObject) (api.ListPassengerAirlineLoyaltiesResponseObject, error) {
	passenger, err := getPassengerByID(ctx, h.queries, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.ListPassengerAirlineLoyalties404Response{}, nil
		}
		return nil, err
	}

	loyalties, err := h.GetAirlineLoyaltiesByPassenger(ctx, passenger.ID)
	if err != nil {
		return nil, err
	}

	return api.ListPassengerAirlineLoyalties200JSONResponse(loyalties), nil
}

// AddLoyaltyMiles adds miles to a passenger's loyalty account
func (h *Handler) AddLoyaltyMiles(ctx context.Context, request api.AddLoyaltyMilesRequestObject) (api.AddLoyaltyMilesResponseObject, error) {
	airline, err := getAirlineBySpec(ctx, h.queries, request.Body.Airline) 
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.AddLoyaltyMiles404Response{}, nil
		}
		return nil, err
	}

	passenger, err := getPassengerByID(ctx, h.queries, int64(request.Body.PassengerId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.AddLoyaltyMiles404Response{}, nil
		}
		return nil, err
	}

	loyalty, err := h.CreateOrUpdateAirlineLoyalty(ctx, airline.ID, passenger.ID, request.Body.Miles)
	if err != nil {
		return nil, err
	}

	return api.AddLoyaltyMiles200JSONResponse(loyalty), nil
}