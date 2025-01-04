import type { paths } from '$lib/airline.openapi'
import type { FlightInstance } from '$lib/types'
import { Type, type TProperties } from '@sinclair/typebox'

const commonProperties = {
	aircraft: Type.Optional(Type.String()),
	notes: Type.Optional(Type.String()),
} satisfies TProperties

export const flightInstanceFromScheduleFormSchema = Type.Object(commonProperties)

export const flightInstanceFromManualInputFormSchema = Type.Object({
	...commonProperties,
	airline: Type.String({ minLength: 2, maxLength: 2 }),
	number: Type.String({ pattern: '^\\d{1,4}$', minLength: 1, maxLength: 4 }),
	originAirport: Type.String({ minLength: 3, maxLength: 3 }),
	destinationAirport: Type.String({ minLength: 3, maxLength: 3 }),
	aircraftType: Type.String(),
	departureDateTime: Type.String(),
	arrivalDateTime: Type.String(),
	published: Type.Boolean(),
})

type FormSchema = (typeof flightInstanceFromScheduleFormSchema)['static'] &
	(typeof flightInstanceFromManualInputFormSchema)['static']

export function existingFlightInstanceToFormData(a: FlightInstance): FormSchema {
	return {
		airline: a.airline.iataCode,
		number: a.number,
		originAirport: a.originAirport.iataCode,
		destinationAirport: a.destinationAirport.iataCode,
		aircraftType: a.aircraftType.icaoCode,
		aircraft: a.aircraft?.registration,
		departureDateTime: a.departureDateTime,
		arrivalDateTime: a.arrivalDateTime,
		notes: a.notes,
		published: a.published,
	}
}

export function formDataToFlightInstanceRequest(
	f: FormSchema,
):
	| paths['/flight-instances']['post']['requestBody']['content']['application/json']
	| paths['/flight-instances/{id}']['patch']['requestBody']['content']['application/json'] {
	return {
		airline: f.airline,
		number: f.number,
		originAirport: f.originAirport,
		destinationAirport: f.destinationAirport,
		aircraftType: f.aircraftType,
		aircraft: f.aircraft,
		departureDateTime: f.departureDateTime,
		arrivalDateTime: f.arrivalDateTime,
		notes: f.notes,
		published: f.published,
	}
}
