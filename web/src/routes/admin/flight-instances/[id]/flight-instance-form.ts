import type { FlightInstance } from '$lib/types'
import { Type } from '@sinclair/typebox'

export const formSchema = Type.Object({
	aircraft: Type.Optional(Type.Integer()),
	notes: Type.Optional(Type.String()),
})

export type FormSchema = typeof formSchema

export function existingFlightInstanceToFormData(a: FlightInstance): FormSchema['static'] {
	return {
		aircraft: a.aircraft?.id,
		notes: a.notes,
	}
}
