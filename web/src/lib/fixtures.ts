import type { Airport } from '$lib/types'

const airports = [
	{
		id: 1,
		iataCode: 'SFO',
		name: 'San Francisco International Airport',
		country: 'US',
		region: 'California',
		point: {
			longitude: -122.375,
			latitude: 37.621,
		},
	},
	{
		id: 2,
		iataCode: 'EWR',
		name: 'Newark Liberty International Airport',
		country: 'US',
		region: 'New Jersey',
		point: {
			longitude: -74.175,
			latitude: 40.692,
		},
	},
	{
		id: 3,
		iataCode: 'LHR',
		name: 'London Heathrow Airport',
		country: 'United Kingdom',
		region: 'England',
		point: {
			longitude: -0.452,
			latitude: 51.471,
		},
	},
	{
		id: 4,
		iataCode: 'AMS',
		name: 'Amsterdam Airport Schiphol',
		country: 'Netherlands',
		region: 'North Holland',
		point: {
			longitude: 4.764,
			latitude: 52.308,
		},
	},
	{
		id: 5,
		iataCode: 'SIN',
		name: 'Singapore Changi Airport',
		country: 'Singapore',
		region: 'Central Region',
		point: {
			longitude: 103.994,
			latitude: 1.35,
		},
	},
] as const

type AirportIATACode = (typeof airports)[number]['iataCode']

export const FIXTURE_AIRPORTS: Record<AirportIATACode, Airport> = Object.fromEntries(
	airports.map((airport) => [airport.iataCode, airport]),
) as Record<AirportIATACode, Airport>
