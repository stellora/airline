package main

import (
	"context"
	"fmt"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/tidwall/geodesic"
)

func getFlight(id int) *api.Flight {
	for _, flight := range flights {
		if flight.Id == id {
			return flight
		}
	}
	return nil
}

func fromDBFlight(a db.FlightsView) api.Flight {
	b := api.Flight{
		Id:        int(a.ID),
		Number:    a.Number,
		Published: a.Published,
	}

	b.OriginAirport = fromDBAirport(db.Airport{
		ID:       a.OriginAirportID,
		IataCode: a.OriginAirportIataCode,
		OadbID:   a.OriginAirportOadbID,
	})
	b.DestinationAirport = fromDBAirport(db.Airport{
		ID:       a.DestinationAirportID,
		IataCode: a.DestinationAirportIataCode,
		OadbID:   a.DestinationAirportOadbID,
	})

	var distanceMeters float64
	geodesic.WGS84.Inverse(b.OriginAirport.Point.Latitude, b.OriginAirport.Point.Longitude, b.DestinationAirport.Point.Latitude, b.DestinationAirport.Point.Longitude, &distanceMeters, nil, nil)
	const metersPerMile = 0.000621371192237334
	b.DistanceMiles = distanceMeters * metersPerMile

	return b
}

func (h *Handler) GetFlight(ctx context.Context, request api.GetFlightRequestObject) (api.GetFlightResponseObject, error) {
	flight := getFlight(request.Id)
	if flight == nil {
		return &api.GetFlight404Response{}, nil
	}
	enrichFlight(flight)
	return api.GetFlight200JSONResponse(*flight), nil
}

func copyFlights(flights []*api.Flight) []api.Flight {
	copies := make([]api.Flight, len(flights))
	for i, flight := range flights {
		copies[i] = *flight
		enrichFlight(&copies[i])
	}
	return copies
}

func enrichFlight(flight *api.Flight) {
	flight.OriginAirport = *getAirport(flight.OriginAirport.Id)
	flight.DestinationAirport = *getAirport(flight.DestinationAirport.Id)

	var distanceMeters float64
	geodesic.WGS84.Inverse(flight.OriginAirport.Point.Latitude, flight.OriginAirport.Point.Longitude, flight.DestinationAirport.Point.Latitude, flight.DestinationAirport.Point.Longitude, &distanceMeters, nil, nil)
	const metersPerMile = 0.000621371192237334
	flight.DistanceMiles = distanceMeters * metersPerMile
}

func (h *Handler) ListFlights(ctx context.Context, request api.ListFlightsRequestObject) (api.ListFlightsResponseObject, error) {
	return api.ListFlights200JSONResponse(copyFlights(flights)), nil
}

func (h *Handler) CreateFlight(ctx context.Context, request api.CreateFlightRequestObject) (api.CreateFlightResponseObject, error) {
	if request.Body.Number == "" {
		return nil, fmt.Errorf("number must not be empty")
	}

	originAirport := getAirportBySpec(request.Body.OriginAirport)
	if originAirport == nil {
		return nil, fmt.Errorf("originAirport not found")
	}
	destinationAirport := getAirportBySpec(request.Body.DestinationAirport)
	if destinationAirport == nil {
		return nil, fmt.Errorf("destinationAirport not found")
	}

	flights = append(flights, &api.Flight{
		Id:                 len(flights) + 1,
		Number:             request.Body.Number,
		OriginAirport:      *originAirport,
		DestinationAirport: *destinationAirport,
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
