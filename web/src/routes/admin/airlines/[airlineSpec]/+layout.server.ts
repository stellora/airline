import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Airline } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/airlines/{airlineSpec}', {
		params: { path: { airlineSpec: params.airlineSpec } },
		fetch,
	})
	const airline = resp.data
	if (!airline) {
		error(resp.response.status, resp.error)
	}
	return {
		airline,
		form: await superValidate(
			existingAirlineToFormData(airline),
			typebox(schema['/airlines/{airlineSpec}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingAirlineToFormData(
	a: Airline,
): Static<(typeof schema)['/airlines/{airlineSpec}']['PATCH']['args']['properties']['body']> {
	return {
		iataCode: a.iataCode,
		name: a.name,
	}
}
