package main

import "github.com/stellora/airline/api-server/api"

var (
	airports []*api.Airport
	flights  []*api.Flight
)

func getAirport(id int) *api.Airport {
	for _, airport := range airports {
		if airport.Id == id {
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

func getFlight(id int) *api.Flight {
	for _, flight := range flights {
		if flight.Id == id {
			return flight
		}
	}
	return nil
}

func init() {
	initialAirports := []string{"SFO", "EWR", "SIN"}
	for _, iataCode := range initialAirports {
		airports = append(airports, &api.Airport{Id: len(airports) + 1, IataCode: iataCode})
	}

	initialFlights := []api.Flight{
		{Number: "UA1", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(3), Published: true},
		{Number: "UA2", OriginAirport: *getAirport(3), DestinationAirport: *getAirport(1), Published: true},
		{Number: "UA2168", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(2), Published: true},
		{Number: "UA1054", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(2), Published: true},
		{Number: "UA2855", OriginAirport: *getAirport(2), DestinationAirport: *getAirport(1), Published: true},
		{Number: "UA598", OriginAirport: *getAirport(2), DestinationAirport: *getAirport(1), Published: true},
	}
	for _, flight := range initialFlights {
		flight.Id = len(flights) + 1
		flights = append(flights, &flight)
	}
}
