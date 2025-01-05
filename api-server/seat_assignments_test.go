package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListSeatAssignmentsForFlightInstance(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe", "Jane Doe")
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)
	insertSeatAssignmentT(t, handler, int64(passengers[0]), flight.Id, "1A")
	insertSeatAssignmentT(t, handler, int64(passengers[1]), flight.Id, "3D")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.ListSeatAssignmentsForFlightInstance(ctx, api.ListSeatAssignmentsForFlightInstanceRequestObject{
			FlightInstanceID: flight.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		got := resp.(api.ListSeatAssignmentsForFlightInstance200JSONResponse)
		assertEqual(t, got, api.ListSeatAssignmentsForFlightInstance200JSONResponse{
			{
				Id:               1,
				ItineraryID:      1,
				PassengerID:      int(passengers[0]),
				FlightInstanceID: flight.Id,
				Seat:             "1A",
			},
			{
				Id:               2,
				ItineraryID:      2,
				PassengerID:      int(passengers[1]),
				FlightInstanceID: flight.Id,
				Seat:             "3D",
			},
		})
	})

	t.Run("flight does not exist", func(t *testing.T) {
		resp, err := handler.ListSeatAssignmentsForFlightInstance(ctx, api.ListSeatAssignmentsForFlightInstanceRequestObject{
			FlightInstanceID: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.ListSeatAssignmentsForFlightInstance404Response{})
	})
}

func TestCreateSeatAssignment(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passenger := insertPassengersWithNamesT(t, handler, "John Doe")[0]
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)

	resp, err := handler.CreateSeatAssignment(ctx, api.CreateSeatAssignmentRequestObject{
		FlightInstanceID: flight.Id,
		Body: &api.CreateSeatAssignmentJSONRequestBody{
			PassengerID: int(passenger),
			Seat:        "1A",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.CreateSeatAssignment201JSONResponse)
	assertEqual(t, got, api.CreateSeatAssignment201JSONResponse{
		Id:               1,
		PassengerID:      int(passenger),
		FlightInstanceID: flight.Id,
		Seat:             "1A",
	})
}

func insertSeatAssignmentT(t *testing.T, handler *Handler, passengerID int64, flightInstanceID int, seat string) int64 {
	// Insert an itinerary for convenience.
	itineraryID := insertItineraryT(t, handler, []int64{int64(flightInstanceID)}, []int64{passengerID})

	resp, err := handler.CreateSeatAssignment(context.Background(), api.CreateSeatAssignmentRequestObject{
		FlightInstanceID: flightInstanceID,
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
