package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListSeatAssignmentsForFlight(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe", "Jane Doe")
	flight := insertFlightT(t, handler, fixtureManualFlight)
	insertSeatAssignmentT(t, handler, int64(passengers[0]), flight.Id, "1A")
	insertSeatAssignmentT(t, handler, int64(passengers[1]), flight.Id, "3D")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.ListSeatAssignmentsForFlight(ctx, api.ListSeatAssignmentsForFlightRequestObject{
			FlightID: flight.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		got := resp.(api.ListSeatAssignmentsForFlight200JSONResponse)
		assertEqual(t, got, api.ListSeatAssignmentsForFlight200JSONResponse{
			{
				Id: 1,
				Itinerary: api.ItinerarySpecs{
					Id:       1,
					RecordID: "TEST00",
				},
				Passenger: api.Passenger{
					Id:   passengers[0],
					Name: "John Doe",
				},
				FlightID: flight.Id,
				Seat:             "1A",
			},
			{
				Id: 2,
				Itinerary: api.ItinerarySpecs{
					Id:       2,
					RecordID: "TEST01",
				},
				Passenger: api.Passenger{
					Id:   passengers[1],
					Name: "Jane Doe",
				},
				FlightID: flight.Id,
				Seat:             "3D",
			},
		})
	})

	t.Run("flight does not exist", func(t *testing.T) {
		resp, err := handler.ListSeatAssignmentsForFlight(ctx, api.ListSeatAssignmentsForFlightRequestObject{
			FlightID: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.ListSeatAssignmentsForFlight404Response{})
	})
}

func TestCreateSeatAssignment(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passenger := insertPassengersWithNamesT(t, handler, "John Doe")[0]
	flight := insertFlightT(t, handler, fixtureManualFlight)
	itineraryID := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{int64(passenger)})

	resp, err := handler.CreateSeatAssignment(ctx, api.CreateSeatAssignmentRequestObject{
		FlightID: flight.Id,
		Body: &api.CreateSeatAssignmentJSONRequestBody{
			ItineraryID: api.ItineraryID(itineraryID),
			PassengerID: int(passenger),
			Seat:        "1A",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.CreateSeatAssignment201JSONResponse)
	assertEqual(t, got, api.CreateSeatAssignment201JSONResponse{
		Id: 1,
		Itinerary: api.ItinerarySpecs{
			Id:       1,
			RecordID: "TEST00",
		},
		Passenger: api.Passenger{
			Id:   int(passenger),
			Name: "John Doe",
		},
		FlightID: flight.Id,
		Seat:             "1A",
	})
}

func insertSeatAssignmentT(t *testing.T, handler *Handler, passengerID int64, flightID int, seat string) int64 {
	// Insert an itinerary for convenience.
	itineraryID := insertItineraryT(t, handler, []int64{int64(flightID)}, []int64{passengerID})

	resp, err := handler.CreateSeatAssignment(context.Background(), api.CreateSeatAssignmentRequestObject{
		FlightID: flightID,
		Body: &api.CreateSeatAssignmentJSONRequestBody{
			ItineraryID: int(itineraryID),
			PassengerID: int(passengerID),
			Seat:        seat,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	created := resp.(api.CreateSeatAssignment201JSONResponse)
	return int64(created.Id)
}
