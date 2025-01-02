package main

import (
	"context"

	"github.com/stellora/airline/api-server/api"
)

// aircraftTypes is a list of many popular aircraft types and their ICAO codes. See
// https://www.icao.int/publications/doc8643/pages/search.aspx.
var aircraftTypes = []api.AircraftType{
	{IcaoCode: "A318", Name: "Airbus A318"},
	{IcaoCode: "A319", Name: "Airbus A319"},
	{IcaoCode: "A320", Name: "Airbus A320"},
	{IcaoCode: "A321", Name: "Airbus A321"},
	{IcaoCode: "A19N", Name: "Airbus A319neo"},
	{IcaoCode: "A20N", Name: "Airbus A320neo"},
	{IcaoCode: "A21N", Name: "Airbus A321neo"},
	{IcaoCode: "A332", Name: "Airbus A330-200"},
	{IcaoCode: "A338", Name: "Airbus A330-800"},
	{IcaoCode: "A339", Name: "Airbus A330-900"},
	{IcaoCode: "A342", Name: "Airbus A340-200"},
	{IcaoCode: "A343", Name: "Airbus A340-300"},
	{IcaoCode: "A345", Name: "Airbus A340-500"},
	{IcaoCode: "A346", Name: "Airbus A340-600"},
	{IcaoCode: "A359", Name: "Airbus A350-900"},
	{IcaoCode: "A35K", Name: "Airbus A350-1000"},
	{IcaoCode: "A388", Name: "Airbus A380-800"},
	{IcaoCode: "B37M", Name: "Boeing 737 MAX 7"},
	{IcaoCode: "B38M", Name: "Boeing 737 MAX 8"},
	{IcaoCode: "B39M", Name: "Boeing 737 MAX 9"},
	{IcaoCode: "B3XM", Name: "Boeing 737 MAX 10"},
	{IcaoCode: "B712", Name: "Boeing 717"},
	{IcaoCode: "B737", Name: "Boeing 737-700"},
	{IcaoCode: "B738", Name: "Boeing 737-800"},
	{IcaoCode: "B739", Name: "Boeing 737-900"},
	{IcaoCode: "B744", Name: "Boeing 747-400"},
	{IcaoCode: "B748", Name: "Boeing 747-8I"},
	{IcaoCode: "B752", Name: "Boeing 757-200"},
	{IcaoCode: "B753", Name: "Boeing 757-300"},
	{IcaoCode: "B762", Name: "Boeing 767-200"},
	{IcaoCode: "B763", Name: "Boeing 767-300"},
	{IcaoCode: "B764", Name: "Boeing 767-400"},
	{IcaoCode: "B772", Name: "Boeing 777-200"},
	{IcaoCode: "B77L", Name: "Boeing 777-200LR"},
	{IcaoCode: "B773", Name: "Boeing 777-300"},
	{IcaoCode: "B77W", Name: "Boeing 777-300ER"},
	{IcaoCode: "B788", Name: "Boeing 787-8"},
	{IcaoCode: "B789", Name: "Boeing 787-9"},
	{IcaoCode: "B78X", Name: "Boeing 787-10"},
	{IcaoCode: "CRJ2", Name: "Bombardier CRJ-200"},
	{IcaoCode: "CRJ7", Name: "Bombardier CRJ-700"},
	{IcaoCode: "CRJ9", Name: "Bombardier CRJ-900"},
	{IcaoCode: "CRJX", Name: "Bombardier CRJ-1000"},
	{IcaoCode: "E170", Name: "Embraer ERJ-170"},
	{IcaoCode: "E75L", Name: "Embraer ERJ-175"},
	{IcaoCode: "E190", Name: "Embraer ERJ-190"},
	{IcaoCode: "E195", Name: "Embraer ERJ-195"},
	{IcaoCode: "BCS1", Name: "Airbus A220-100"},
	{IcaoCode: "BCS3", Name: "Airbus A220-300"},
	{IcaoCode: "DH8D", Name: "Bombardier Q400"},
	{IcaoCode: "AT72", Name: "ATR 72"},
	{IcaoCode: "AT76", Name: "ATR 72-600"},
	{IcaoCode: "AT75", Name: "ATR 72-500"},
	{IcaoCode: "AT45", Name: "ATR 42-500"},
}

func (h *Handler) ListAircraftTypes(ctx context.Context, request api.ListAircraftTypesRequestObject) (api.ListAircraftTypesResponseObject, error) {
	return api.ListAircraftTypes200JSONResponse(aircraftTypes), nil
}
