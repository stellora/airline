package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/stellora/airline/api-server/api"
)

var (
	airports = []api.Airport{
		{Title: "Silverware"},
		{Title: "Cookware"},
		{Title: "Vegetables"},
	}
)

func getAirport(id string) *api.Airport {
	for i := range airports {
		if airports[i].Id == id {
			return &airports[i]
		}
	}
	return nil
}

func init() {
	for i := range airports {
		airports[i].Id = strconv.Itoa(i + 1)
	}
}

func (h *Handler) GetAirport(ctx context.Context, request api.GetAirportRequestObject) (api.GetAirportResponseObject, error) {
	airport := getAirport(request.Id)
	if airport == nil {
		return &api.GetAirport404Response{}, nil
	}
	return api.GetAirport200JSONResponse(*airport), nil
}

func (h *Handler) ListAirports(ctx context.Context, request api.ListAirportsRequestObject) (api.ListAirportsResponseObject, error) {
	return api.ListAirports200JSONResponse(airports), nil
}

func (h *Handler) CreateAirport(ctx context.Context, request api.CreateAirportRequestObject) (api.CreateAirportResponseObject, error) {
	title := request.Body.Title
	if title == "" {
		return nil, fmt.Errorf("title must not be empty")
	}

	for _, airport := range airports {
		if airport.Title == title {
			return nil, fmt.Errorf("title must be unique across all airports")
		}
	}

	newAirport := api.Airport{
		Id:    strconv.Itoa(len(airports) + 1),
		Title: title,
	}
	airports = append(airports, newAirport)

	return api.CreateAirport201Response{}, nil
}

func (h *Handler) DeleteAirport(ctx context.Context, request api.DeleteAirportRequestObject) (api.DeleteAirportResponseObject, error) {
	// Find and remove the airport
	newAirports := []api.Airport{}
	found := false
	for _, airport := range airports {
		if airport.Id != request.Id {
			newAirports = append(newAirports, airport)
		} else {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("airport not found")
	}

	// Remove all flight memberships for this airport
	newMemberships := []flightAirportMembership{}
	for _, membership := range flightAirportMemberships {
		if membership.airport != request.Id {
			newMemberships = append(newMemberships, membership)
		}
	}

	airports = newAirports
	flightAirportMemberships = newMemberships

	return api.DeleteAirport204Response{}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
