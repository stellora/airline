import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Aircraft } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/aircraft/{aircraftSpec}', {
		params: { path: { aircraftSpec: params.aircraftSpec } },
		fetch,
	})
	const aircraft = resp.data
	if (!aircraft) {
		error(resp.response.status, resp.error)
	}
	return {
		aircraft,
		form: await superValidate(existingAircraftToFormData(aircraft), typebox(schema['/aircraft/{aircraftSpec}']['PATCH']['args']['properties']['body'])),
	}
}

function existingAircraftToFormData(
	a: Aircraft,
): Static<(typeof schema)['/aircraft/{aircraftSpec}']['PATCH']['args']['properties']['body']> {
	return {
		registration: a.registration,
		aircraftType: a.aircraftType,
		airline: a.airline.iataCode,
	}
}
