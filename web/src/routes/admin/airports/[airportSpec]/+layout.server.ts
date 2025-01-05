import { apiClient } from '$lib/api'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { existingAirportToFormData, formSchema } from '../airport-form'
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
		form: await superValidate(existingAirportToFormData(airport), typebox(formSchema)),
	}
}
