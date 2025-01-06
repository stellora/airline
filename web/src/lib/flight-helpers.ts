import * as d3 from 'd3'
import type { Feature } from 'geojson'
import type { Airport, DaysOfWeek, FlightInstance, Schedule } from './types'

export function flightTitle(
	flight: Pick<Schedule, 'airline' | 'number' | 'originAirport' | 'destinationAirport'>,
): string {
	return `${flight.airline.iataCode} ${flight.number} ${flightRoute(flight.originAirport, flight.destinationAirport)}`
}

export function flightInstanceTitle(flight: FlightInstance): string {
	return `${flightTitle(flight)} on ${flight.scheduleInstanceDate}`
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

const DAYS_OF_WEEK_NARROW = ['Su', 'M', 'Tu', 'W', 'Th', 'F', 'Sa']
export function formatDaysOfWeek(daysOfWeek: DaysOfWeek): string {
	return daysOfWeek.map((d) => DAYS_OF_WEEK_NARROW[d]).join('-')
}
