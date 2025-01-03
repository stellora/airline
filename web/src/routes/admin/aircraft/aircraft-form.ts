import type { paths } from '$lib/airline.openapi'
import type { Aircraft } from '$lib/types'
import { Type } from '@sinclair/typebox'

// TODO!(sqs): use codegen
export const formSchema = Type.Object({
	registration: Type.String(),
	aircraftType: Type.String(),
	airline: Type.String(),
})

export type FormSchema = typeof formSchema

export function existingAircraftToFormData(a: Aircraft): FormSchema['static'] {
	return {
		registration: a.registration,
		aircraftType: a.aircraftType,
		airline: a.airline.iataCode,
	}
}

export function formDataToAircraftRequest(
	f: FormSchema['static'],
):
	| paths['/aircraft']['post']['requestBody']['content']['application/json']
	| paths['/aircraft']['post']['requestBody']['content']['application/json'] {
	return {
		registration: f.registration,
		aircraftType: f.aircraftType,
		airline: f.airline,
	}
}
