import { apiClient } from '$lib/api'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { existingAircraftToFormData, formSchema } from '../aircraft-form'
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
		form: await superValidate(existingAircraftToFormData(aircraft), typebox(formSchema)),
	}
}
