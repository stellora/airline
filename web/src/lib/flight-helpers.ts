import type { Flight } from './types'

export function flightTitle(flight: Flight): string {
	return `${flight.number} ${flight.originAirport.iataCode}–${flight.destinationAirport.iataCode}`
}
