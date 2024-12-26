package main

import (
	"context"
	"fmt"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) GetFlight(ctx context.Context, request api.GetFlightRequestObject) (api.GetFlightResponseObject, error) {
	flight := getFlight(request.Id)
	if flight == nil {
		return &api.GetFlight404Response{}, nil
	}
	populateFlightAirports(flight)
	return api.GetFlight200JSONResponse(*flight), nil
}

func populateFlightAirports(flight *api.Flight) {
	flight.OriginAirport = *getAirport(flight.OriginAirport.Id)
	flight.DestinationAirport = *getAirport(flight.DestinationAirport.Id)
}

func copyFlights(flights []*api.Flight) []api.Flight {
	copies := make([]api.Flight, len(flights))
	for i, flight := range flights {
		copies[i] = *flight
	}
	return copies
}

func (h *Handler) ListFlights(ctx context.Context, request api.ListFlightsRequestObject) (api.ListFlightsResponseObject, error) {
	flights := copyFlights(flights)
	for i := range flights {
		populateFlightAirports(&flights[i])
	}
	return api.ListFlights200JSONResponse(flights), nil
}

func (h *Handler) CreateFlight(ctx context.Context, request api.CreateFlightRequestObject) (api.CreateFlightResponseObject, error) {
	if request.Body.Number == "" {
		return nil, fmt.Errorf("number must not be empty")
	}
	if request.Body.OriginAirport == 0 {
		return nil, fmt.Errorf("originAirport must not be empty")
	}
	if request.Body.DestinationAirport == 0 {
		return nil, fmt.Errorf("destinationAirport must not be empty")
	}

	flights = append(flights, &api.Flight{
		Id:                 len(flights) + 1,
		Number:             request.Body.Number,
		OriginAirport:      api.Airport{Id: request.Body.OriginAirport},
		DestinationAirport: api.Airport{Id: request.Body.DestinationAirport},
		Published:          request.Body.Published != nil && *request.Body.Published,
	})
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
	flights = []*api.Flight{}
	return api.DeleteAllFlights204Response{}, nil
}
