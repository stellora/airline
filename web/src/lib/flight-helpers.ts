import type { Airport, Flight } from './types'

export function flightTitle(flight: Flight): string {
	return `${flight.number} ${flightRoute(flight.originAirport, flight.destinationAirport)}`
}

export function flightRoute(
	origin: Pick<Airport, 'iataCode'>,
	destination: Pick<Airport, 'iataCode'>
): string {
	return `${origin.iataCode}–${destination.iataCode}`
}
