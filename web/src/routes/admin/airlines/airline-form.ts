import type { paths } from '$lib/airline.openapi'
import { Type } from '@sinclair/typebox'

// TODO!(sqs): use codegen
export const formSchema = Type.Object({
	iataCode: Type.String({ minLength: 2, maxLength: 2 }),
	name: Type.String({}),
})

export type FormSchema = typeof formSchema

export function formDataToAirlineRequest(
	f: FormSchema['static'],
):
	| paths['/airlines']['post']['requestBody']['content']['application/json']
	| paths['/airlines/{airlineSpec}']['patch']['requestBody']['content']['application/json'] {
	return {
		iataCode: f.iataCode,
		name: f.name,
	}
}
