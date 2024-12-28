package main

import "github.com/stellora/airline/api-server/api"

var (
	airports []*api.Airport
	flights  []*api.Flight
)

// TODO!(sqs): add sample data
func init() {
	// initialAirports := []string{"SFO", "EWR", "SIN"}
	// for _, iataCode := range initialAirports {
	// 	airports = append(airports, &api.Airport{Id: len(airports) + 1, IataCode: iataCode})
	// }
	//
	// initialFlights := []api.Flight{
	// 	{Number: "UA1", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(3), Published: true},
	// 	{Number: "UA2", OriginAirport: *getAirport(3), DestinationAirport: *getAirport(1), Published: true},
	// 	{Number: "UA2168", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(2), Published: true},
	// 	{Number: "UA1054", OriginAirport: *getAirport(1), DestinationAirport: *getAirport(2), Published: true},
	// 	{Number: "UA2855", OriginAirport: *getAirport(2), DestinationAirport: *getAirport(1), Published: true},
	// 	{Number: "UA598", OriginAirport: *getAirport(2), DestinationAirport: *getAirport(1), Published: true},
	// }
	// for _, flight := range initialFlights {
	// 	flight.Id = len(flights) + 1
	// 	flights = append(flights, &flight)
	// }
}
