package main

import (
	"context"
	"fmt"

	"github.com/stellora/airline/api-server/api"
)

func (h *Handler) GetAirport(ctx context.Context, request api.GetAirportRequestObject) (api.GetAirportResponseObject, error) {
	airport := getAirport(request.Id)
	if airport == nil {
		return &api.GetAirport404Response{}, nil
	}
	return api.GetAirport200JSONResponse(*airport), nil
}

func copyAirports(airports []*api.Airport) []api.Airport {
	copies := make([]api.Airport, len(airports))
	for i, airport := range airports {
		copies[i] = *airport
	}
	return copies
}

func (h *Handler) ListAirports(ctx context.Context, request api.ListAirportsRequestObject) (api.ListAirportsResponseObject, error) {
	return api.ListAirports200JSONResponse(copyAirports(airports)), nil
}

func (h *Handler) CreateAirport(ctx context.Context, request api.CreateAirportRequestObject) (api.CreateAirportResponseObject, error) {
	IataCode := request.Body.IataCode
	if IataCode == "" {
		return nil, fmt.Errorf("iataCode must not be empty")
	}

	for _, airport := range airports {
		if airport.IataCode == IataCode {
			return nil, fmt.Errorf("iataCode must be unique across all airports")
		}
	}

	airports = append(airports, &api.Airport{
		Id:       len(airports) + 1,
		IataCode: IataCode,
	})
	return api.CreateAirport201Response{}, nil
}

func (h *Handler) UpdateAirport(ctx context.Context, request api.UpdateAirportRequestObject) (api.UpdateAirportResponseObject, error) {
	airport := getAirport(request.Id)
	if airport == nil {
		return &api.UpdateAirport404Response{}, nil
	}

	if request.Body.IataCode != nil {
		airport.IataCode = *request.Body.IataCode
	}
	return api.UpdateAirport204Response{}, nil
}

func (h *Handler) DeleteAirport(ctx context.Context, request api.DeleteAirportRequestObject) (api.DeleteAirportResponseObject, error) {
	// Find and remove the airport
	for i, airport := range airports {
		if airport.Id == request.Id {
			airports = append(airports[:i], airports[i+1:]...)
			break
		}
	}
	return api.DeleteAirport204Response{}, nil
}

func (h *Handler) DeleteAllAirports(ctx context.Context, request api.DeleteAllAirportsRequestObject) (api.DeleteAllAirportsResponseObject, error) {
	airports = []*api.Airport{}
	return api.DeleteAllAirports204Response{}, nil
}

func newAirportSpec(id int, iataCode string) api.AirportSpec {
	var spec api.AirportSpec
	if id != 0 {
		spec.FromAirportSpec0(id)
	} else {
		spec.FromAirportSpec1(iataCode)
	}
	return spec
}
