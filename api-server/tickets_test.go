package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

var (
	fixtureTicket1 = api.Ticket{
		Id:        1,
		Number:    "0160000000000",
		Passenger: fixturePassenger1,
		FareBasis: "Y",
	}
	fixtureTicket2 = api.Ticket{
		Id:        2,
		Number:    "0160000000001",
		Passenger: fixturePassenger2,
		FareBasis: "Y",
	}
)

func TestGetTicket(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passenger := insertPassengersWithNamesT(t, handler, "John Doe")[0]
	flight := insertFlightT(t, handler, fixtureManualFlight)
	itinerary := insertItineraryT(t, handler, []int64{int64(flight.Id)}, []int64{passenger})
	ticketNumber, _ := insertTicketT(t, handler, "1", itinerary, passenger)

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetTicket(ctx, api.GetTicketRequestObject{
			TicketNumber: ticketNumber,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.GetTicket200JSONResponse(fixtureTicket1))
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetTicket(ctx, api.GetTicketRequestObject{
			TicketNumber: "1234567890123",
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetTicket404Response{})
	})
}

func TestListTickets(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	passengers := insertPassengersWithNamesT(t, handler, "John Doe", "Alice Zhao")
	flight := insertFlightT(t, handler, fixtureManualFlight)
	itinerary := insertItineraryT(t, handler, []int64{int64(flight.Id)}, passengers)
	ticketNumber1, ticketID1 := insertTicketT(t, handler, "1", itinerary, passengers[0])
	ticketNumber2, ticketID2 := insertTicketT(t, handler, "2", itinerary, passengers[1])

	resp, err := handler.ListTickets(ctx, api.ListTicketsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, resp, api.ListTickets200JSONResponse{
		api.Ticket{
			Id:     int(ticketID1),
			Number: ticketNumber1,
			Itinerary: api.ItinerarySpecs{
				Id:       int(itinerary),
				RecordID: api.RecordLocator("TEST01"),
			},
			Passenger: fixturePassenger1,
			FareBasis: "Y",
		},
		api.Ticket{
			Id:     int(ticketID2),
			Number: ticketNumber2,
			Itinerary: api.ItinerarySpecs{
				Id:       int(itinerary),
				RecordID: api.RecordLocator("TEST01"),
			},
			Passenger: fixturePassenger2,
			FareBasis: "Y",
		},
	})
}

func insertTicketT(t *testing.T, handler *Handler, suffix string, itineraryID, passengerID int64) (ticketNumber string, id int64) {
	t.Helper()
	ctx := context.Background()
	resp, err := handler.queries.CreateTicket(ctx, db.CreateTicketParams{
		Number:      fmt.Sprintf("016000000000%s", suffix),
		ItineraryID: itineraryID,
		PassengerID: passengerID,
		FareBasis:   "Y",
	})
	if err != nil {
		t.Fatal(err)
	}
	return resp.Number, resp.ID
}
