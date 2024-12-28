package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
)

func getAirport(id int) *api.Airport {
	for _, airport := range airports {
		if airport.Id == id {
			enrichAirport(airport)
			return airport
		}
	}
	return nil
}

func getAirportBySpec(spec api.AirportSpec) *api.Airport {
	if id, err := spec.AsAirportSpec0(); err == nil {
		return getAirport(id)
	}
	if iataCode, err := spec.AsAirportSpec1(); err == nil {
		for _, airport := range airports {
			if airport.IataCode == iataCode {
				return airport
			}
		}
	}
	return nil
}

func fromDBAirport(a db.Airport) api.Airport {
	b := api.Airport{
		Id:       int(a.ID),
		IataCode: a.IataCode,
	}
	if a.OadbID.Valid {
		info := extdata.Airports.AirportByOAID(int(a.OadbID.Int64))
		b.Name = info.Airport.Name
		b.Country = string(info.Country.Code)
		b.Region = info.Region.Name
		b.Point = api.Point{
			Latitude:  info.Airport.LatitudeDeg,
			Longitude: info.Airport.LongitudeDeg,
		}
	}
	// TODO(sqs): handle case where there is no oadb_id or the OA database has no airport with the ID
	return b
}

func (h *Handler) GetAirport(ctx context.Context, request api.GetAirportRequestObject) (api.GetAirportResponseObject, error) {
	airport, err := h.queries.GetAirport(ctx, int64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &api.GetAirport404Response{}, nil
		}
	}
	return api.GetAirport200JSONResponse(fromDBAirport(airport)), nil
}

func copyAirports(airports []*api.Airport) []api.Airport {
	copies := make([]api.Airport, len(airports))
	for i, airport := range airports {
		copies[i] = *airport
		enrichAirport(&copies[i])
	}
	return copies
}

func enrichAirport(airport *api.Airport) {
	airportData := lookupAirport(airport.IataCode)
	if airportData == nil {
		return
	}
	airport.Name = airportData.Name
	airport.Point = api.Point{Latitude: airportData.LatitudeDeg, Longitude: airportData.LongitudeDeg}
	airport.Country = string(airportData.ISOCountry)
	airport.Region = string(airportData.Municipality)
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

func lookupAirport(iataCode string) *extdata.Airport {
	iataCode = strings.ToUpper(iataCode)
	for _, airport := range extdata.Airports.Airports {
		if airport.IATACode == iataCode {
			return &airport
		}
	}
	return nil
}
