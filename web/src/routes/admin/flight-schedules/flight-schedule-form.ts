import type { FlightSchedule } from '$lib/types'
import { Type } from '@sinclair/typebox'

export const formSchema = Type.Object({
	airline: Type.String({ minLength: 2, maxLength: 2 }),
	number: Type.String({ pattern: '^\\d{1,4}$', minLength: 1, maxLength: 4 }),
	originAirport: Type.String({ minLength: 3, maxLength: 3 }),
	destinationAirport: Type.String({ minLength: 3, maxLength: 3 }),
	aircraftType: Type.String(),
	startEndDate: Type.Object({
		start: Type.String(),
		end: Type.String(),
	}),
	daysOfWeek: Type.Array(Type.Integer()),
	departureTime: Type.String(),
	arrivalTime: Type.String(),
	published: Type.Boolean(),
})

export type FormSchema = typeof formSchema

export function existingFlightScheduleToFormData(a: FlightSchedule): FormSchema['static'] {
	return {
		airline: a.airline.iataCode,
		number: a.number,
		originAirport: a.originAirport.iataCode,
		destinationAirport: a.destinationAirport.iataCode,
		aircraftType: a.aircraftType.icaoCode,
		startEndDate: {
			start: a.startDate,
			end: a.endDate,
		},
		daysOfWeek: a.daysOfWeek,
		departureTime: a.departureTime,
		arrivalTime: a.arrivalTime,
		published: a.published,
	}
}
