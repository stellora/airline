import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Passenger } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/passengers/{id}', {
		params: { path: { id: Number.parseInt(params.id) } },
		fetch,
	})
	const passenger = resp.data
	if (!passenger) {
		error(resp.response.status, resp.error)
	}
	return {
		passenger,
		form: await superValidate(
			existingPassengerToFormData(passenger),
			typebox(schema['/passengers/{id}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingPassengerToFormData(
	a: Passenger,
): Static<(typeof schema)['/passengers/{id}']['PATCH']['args']['properties']['body']> {
	return {
		name: a.name,
	}
}
