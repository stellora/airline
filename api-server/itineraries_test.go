package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetItinerary(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passenger := insertPassengersWithNamesT(t, handler, "John Doe")[0]
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)
	itinerary := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{int64(passenger)})

	t.Run("exists", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetItinerary(ctx, api.GetItineraryRequestObject{
				ItinerarySpec: api.NewItinerarySpec(int(itinerary), ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			got := resp.(api.GetItinerary200JSONResponse)
			assertEqual(t, got.Id, int(itinerary))
			assertEqual(t, len(got.Flights), 1)
			assertEqual(t, len(got.Passengers), 1)
		})

		t.Run("by record locator", func(t *testing.T) {
			resp, err := handler.GetItinerary(ctx, api.GetItineraryRequestObject{
				ItinerarySpec: api.NewItinerarySpec(0, "ABC123"),
			})
			if err != nil {
				t.Fatal(err)
			}
			got := resp.(api.GetItinerary200JSONResponse)
			assertEqual(t, got.Id, int(itinerary))
		})
	})

	t.Run("does not exist", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetItinerary(ctx, api.GetItineraryRequestObject{
				ItinerarySpec: api.NewItinerarySpec(999, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetItinerary404Response{})
		})

		t.Run("by record locator", func(t *testing.T) {
			resp, err := handler.GetItinerary(ctx, api.GetItineraryRequestObject{
				ItinerarySpec: api.NewItinerarySpec(0, "NOTFND"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetItinerary404Response{})
		})
	})
}

func TestListItineraries(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe", "Jane Doe")
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)
	itinerary1 := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{int64(passengers[0])})
	itinerary2 := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{int64(passengers[1])})

	resp, err := handler.ListItineraries(ctx, api.ListItinerariesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.ListItineraries200JSONResponse)
	assertEqual(t, len(got), 2)
	assertEqual(t, got[0].Id, int(itinerary1))
	assertEqual(t, got[1].Id, int(itinerary2))
}

func TestCreateItinerary(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe")
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)

	resp, err := handler.CreateItinerary(ctx, api.CreateItineraryRequestObject{
		Body: &api.CreateItineraryJSONRequestBody{
			FlightInstanceIDs: []int{int(flight.Id)},
			PassengerIDs:      []int{int(passengers[0])},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.CreateItinerary201JSONResponse)
	assertEqual(t, len(got.Flights), 1)
	assertEqual(t, len(got.Passengers), 1)
	assertEqual(t, got.Flights[0].Id, int(flight.Id))
	assertEqual(t, got.Passengers[0].Id, int(passengers[0]))
}

func TestDeleteItinerary(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe")
	flight := insertFlightInstanceT(t, handler, fixtureManualFlightInstance)
	itinerary := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{int64(passengers[0])})

	resp, err := handler.DeleteItinerary(ctx, api.DeleteItineraryRequestObject{
		ItinerarySpec: api.NewItinerarySpec(int(itinerary), ""),
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.DeleteItinerary204Response{})
	checkItineraryCount(t, handler, 0)
}

func checkItineraryCount(t *testing.T, handler *Handler, want int) {
	resp, err := handler.ListItineraries(context.Background(), api.ListItinerariesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	got := resp.(api.ListItineraries200JSONResponse)
	if len(got) != want {
		t.Errorf("got %d itineraries, want %d", len(got), want)
	}
}

func insertItineraryT(t *testing.T, handler *Handler, flightIDs []int64, passengerIDs []int64) int64 {
	var flightInstanceIds []int
	var passengerIds []int
	for _, id := range flightIDs {
		flightInstanceIds = append(flightInstanceIds, int(id))
	}
	for _, id := range passengerIDs {
		passengerIds = append(passengerIds, int(id))
	}

	resp, err := handler.CreateItinerary(context.Background(), api.CreateItineraryRequestObject{
		Body: &api.CreateItineraryJSONRequestBody{
			FlightInstanceIDs: flightInstanceIds,
			PassengerIDs:      passengerIds,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	created := resp.(api.CreateItinerary201JSONResponse)
	return int64(created.Id)
}
