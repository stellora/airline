import type { paths } from '$lib/airline.openapi'
import type { Airport } from '$lib/types'
import { Type } from '@sinclair/typebox'

// TODO!(sqs): use codegen
export const formSchema = Type.Object({
	iataCode: Type.String({ minLength: 3, maxLength: 3 }),
})

export type FormSchema = typeof formSchema

export function existingAirportToFormData(a: Airport): FormSchema['static'] {
	return {
		iataCode: a.iataCode,
	}
}

export function formDataToAirportRequest(
	f: FormSchema['static'],
):
	| paths['/airports']['post']['requestBody']['content']['application/json']
	| paths['/airports/{airportSpec}']['patch']['requestBody']['content']['application/json'] {
	return {
		iataCode: f.iataCode,
	}
}
