import type { FlightInstance } from '$lib/types'
import { Type, type TProperties } from '@sinclair/typebox'

const commonProperties: TProperties = {
	aircraft: Type.Optional(Type.Integer()),
	notes: Type.Optional(Type.String()),
}

export const flightInstanceFromScheduleFormSchema = Type.Object(commonProperties)

export const flightInstanceFromManualInputFormSchema = Type.Object({
	...commonProperties,
	airline: Type.String({ minLength: 2, maxLength: 2 }),
	number: Type.String({ pattern: '^\\d{1,4}$', minLength: 1, maxLength: 4 }),
	originAirport: Type.String({ minLength: 3, maxLength: 3 }),
	destinationAirport: Type.String({ minLength: 3, maxLength: 3 }),
	aircraftType: Type.String(),
	departureDateTime: Type.Date(),
	arrivalDateTime: Type.Date(),
	published: Type.Boolean(),
})

export function existingFlightInstanceToFormData(
	a: FlightInstance,
): (typeof flightInstanceFromScheduleFormSchema)['static'] &
	(typeof flightInstanceFromManualInputFormSchema)['static'] {
	return {
		airline: a.airline.iataCode,
		number: a.number,
		originAirport: a.originAirport.iataCode,
		destinationAirport: a.destinationAirport.iataCode,
		aircraftType: a.aircraftType.icaoCode,
		aircraft: a.aircraft?.id,
		departureDateTime: new Date(a.departureDateTime),
		arrivalDateTime: new Date(a.arrivalDateTime),
		notes: a.notes,
		published: a.published,
	}
}
