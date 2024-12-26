package main

import (
	"context"
	"fmt"

	"github.com/stellora/airline/api-server/api"
)

var (
	flights = []api.Flight{
		{Title: "Fork"},
		{Title: "Spoon"},
		{Title: "Knife"},
		{Title: "Cast-Iron Pan"},
		{Title: "Baking Sheet"},
		{Title: "Cutting Board"},
		{Title: "Tomato"},
		{Title: "Zucchini"},
		{Title: "Avocado"},
	}
)

func getFlight(id int) *api.Flight {
	for i := range flights {
		if flights[i].Id == id {
			return &flights[i]
		}
	}
	return nil
}

func init() {
	for i := range flights {
		flights[i].Id = i + 1
	}
}

func (h *Handler) GetFlight(ctx context.Context, request api.GetFlightRequestObject) (api.GetFlightResponseObject, error) {
	flight := getFlight(request.Id)
	if flight == nil {
		return &api.GetFlight404Response{}, nil
	}
	populateFlightAirports(flight)
	return api.GetFlight200JSONResponse(*flight), nil
}

func populateFlightAirports(flight *api.Flight) {
	airports := []api.Airport{}
	for _, membership := range flightAirportMemberships {
		if membership.flight == flight.Id {
			if airport := getAirport(membership.airport); airport != nil {
				airports = append(airports, *airport)
			}
		}
	}
	flight.Airports = &airports
}

func (h *Handler) ListFlights(ctx context.Context, request api.ListFlightsRequestObject) (api.ListFlightsResponseObject, error) {
	flightsWithAirports := flights
	for i := range flightsWithAirports {
		populateFlightAirports(&flightsWithAirports[i])
	}
	return api.ListFlights200JSONResponse(flightsWithAirports), nil
}

func (h *Handler) ListFlightsByAirport(ctx context.Context, request api.ListFlightsByAirportRequestObject) (api.ListFlightsByAirportResponseObject, error) {
	airport := request.AirportId

	flightsInAirport := []api.Flight{}
	for _, flight := range flights {
		for _, membership := range flightAirportMemberships {
			if membership.flight == flight.Id && membership.airport == airport {
				flightsInAirport = append(flightsInAirport, flight)
				break
			}
		}
	}

	flightsNotInAirport := []api.Flight{}
	for _, flight := range flights {
		inAirport := false
		for _, membership := range flightAirportMemberships {
			if membership.flight == flight.Id && membership.airport == airport {
				inAirport = true
				break
			}
		}
		if !inAirport {
			flightsNotInAirport = append(flightsNotInAirport, flight)
		}
	}

	return api.ListFlightsByAirport200JSONResponse{
		FlightsInAirport:    flightsInAirport,
		FlightsNotInAirport: flightsNotInAirport,
	}, nil
}

func (h *Handler) CreateFlight(ctx context.Context, request api.CreateFlightRequestObject) (api.CreateFlightResponseObject, error) {
	if request.Body.Title == "" {
		return nil, fmt.Errorf("title must not be empty")
	}

	for _, flight := range flights {
		if flight.Title == request.Body.Title {
			return nil, fmt.Errorf("title must be unique across all flights")
		}
	}

	newFlight := api.Flight{
		Id:        len(flights) + 1,
		Title:     request.Body.Title,
		Published: false,
	}
	flights = append(flights, newFlight)

	return api.CreateFlight201Response{}, nil
}

func (h *Handler) UpdateFlight(ctx context.Context, request api.UpdateFlightRequestObject) (api.UpdateFlightResponseObject, error) {
	flight := getFlight(request.Id)
	if flight == nil {
		return &api.UpdateFlight404Response{}, nil
	}

	if request.Body.Number != nil {
		flight.Number = *request.Body.Number
	}
	if request.Body.OriginAirport != nil {
		flight.OriginAirport = api.Airport{Id: *request.Body.OriginAirport}
	}
	if request.Body.DestinationAirport != nil {
		flight.DestinationAirport = api.Airport{Id: *request.Body.DestinationAirport}
	}
	if request.Body.Published != nil {
		flight.Published = *request.Body.Published
	}
	return api.UpdateFlight204Response{}, nil
}

func (h *Handler) DeleteFlight(ctx context.Context, request api.DeleteFlightRequestObject) (api.DeleteFlightResponseObject, error) {
	// Find and remove the flight
	for i, flight := range flights {
		if flight.Id == request.Id {
			flights = append(flights[:i], flights[i+1:]...)
			break
		}
	}
	return api.DeleteFlight204Response{}, nil
}

func (h *Handler) DeleteAllFlights(ctx context.Context, request api.DeleteAllFlightsRequestObject) (api.DeleteAllFlightsResponseObject, error) {
	flights = []api.Flight{}
	return api.DeleteAllFlights204Response{}, nil
}
