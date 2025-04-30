package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"

	"github.com/stellora/airline/api-server/db"
)

func createSampleLoyaltyPrograms(ctx context.Context, handler *Handler) error {
	log.Println("Creating sample loyalty programs...")
	// Get all airlines
	airlines, err := handler.queries.ListAirlines(ctx)
	if err != nil {
		return err
	}

	// Get passengers
	passengers, err := handler.queries.ListPassengers(ctx)
	if err != nil {
		return err
	}

	if len(passengers) == 0 {
		// Create a few sample passengers if none exist
		passengerNames := []string{"John Smith", "Jane Doe", "Robert Johnson"}
		for _, name := range passengerNames {
			_, err := handler.queries.CreatePassenger(ctx, name)
			if err != nil {
				return err
			}
		}
		
		// Fetch the newly created passengers
		passengers, err = handler.queries.ListPassengers(ctx)
		if err != nil {
			return err
		}
	}

	// Create loyalty programs for each airline and passenger
	for _, airline := range airlines {
		for _, passenger := range passengers {
			// Check if loyalty program already exists
			_, err := handler.queries.GetAirlineLoyalty(ctx, db.GetAirlineLoyaltyParams{
				AirlineID:   airline.ID,
				PassengerID: passenger.ID,
			})
			
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					// Create a random mileage balance (10k to 100k miles)
					mileageBalance := int64(rand.Intn(90000) + 10000)
					
					_, err := handler.queries.CreateAirlineLoyalty(ctx, db.CreateAirlineLoyaltyParams{
						AirlineID:      airline.ID,
						PassengerID:    passenger.ID,
						MileageBalance: mileageBalance,
					})
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
		}
	}

	// Update flights to include mileage rewards based on distance
	flights, err := handler.queries.ListFlights(ctx)
	if err != nil {
		return err
	}

	for _, flight := range flights {
		// Check if the flight already has a mileage reward set
		if flight.MileageReward == 0 {
			// Calculate mileage reward (approximately the distance in miles)
			originAirport, err := handler.queries.GetAirport(ctx, flight.OriginAirportID)
			if err != nil {
				return err
			}
			destinationAirport, err := handler.queries.GetAirport(ctx, flight.DestinationAirportID)
			if err != nil {
				return err
			}
			
			mileageReward := calculateMileageReward(originAirport, destinationAirport)
			
			// Update the flight with the mileage reward
			_, err = handler.queries.UpdateFlight(ctx, db.UpdateFlightParams{
				ID:            flight.ID,
				MileageReward: sql.NullInt64{Valid: true, Int64: mileageReward},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}