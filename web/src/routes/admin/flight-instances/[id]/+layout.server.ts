import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { flightInstanceTitle } from '$lib/flight-helpers'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'
import {
	existingFlightInstanceToFormData,
	flightInstanceFromScheduleFormSchema,
} from './flight-instance-form'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flight-instances/{id}', {
		params: { path: { id } },
		fetch,
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Flight instance not found')
	}
	const flightInstance = resp.data
	return {
		...(await breadcrumbEntry(parent, {
			url: route('/admin/flight-instances/[id]', { params: { id: flightInstance.id.toString() } }),
			title: flightInstanceTitle(flightInstance),
		})),
		flightInstance,
		form: await superValidate(
			existingFlightInstanceToFormData(flightInstance),
			typebox(flightInstanceFromScheduleFormSchema),
		),
	}
}
