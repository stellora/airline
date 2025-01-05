import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Airport } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/airports/{airportSpec}', {
		params: { path: { airportSpec: params.airportSpec } },
		fetch,
	})
	const airport = resp.data
	if (!airport) {
		error(resp.response.status, resp.error)
	}
	return {
		airport,
		form: await superValidate(
			existingAirportToFormData(airport),
			typebox(schema['/airports/{airportSpec}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingAirportToFormData(
	a: Airport,
): Static<(typeof schema)['/airports/{airportSpec}']['PATCH']['args']['properties']['body']> {
	return {
		iataCode: a.iataCode,
	}
}
