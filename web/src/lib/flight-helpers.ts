import * as d3 from 'd3'
import type { Feature } from 'geojson'
import type { Airport, FlightSchedule } from './types'

export function flightTitle(flight: FlightSchedule): string {
	return `${flight.airline.iataCode}${flight.number} ${flightRoute(flight.originAirport, flight.destinationAirport)}`
}

export function flightRoute(
	origin: Pick<Airport, 'iataCode'>,
	destination: Pick<Airport, 'iataCode'>,
): string {
	return `${origin.iataCode}–${destination.iataCode}`
}

export function geoDistanceMiles(line: Feature): number {
	const RADIUS_MILES_PER_RADIAN = 3958.8
	return d3.geoLength(line) * RADIUS_MILES_PER_RADIAN
}
