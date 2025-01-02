import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { flightTitle } from '$lib/flight-helpers'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { existingFlightScheduleToFormData, formSchema } from '../flight-schedule-form'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flight-schedules/{id}', {
		params: { path: { id } },
		fetch,
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Flight schedule not found')
	}
	const flightSchedule = resp.data
	return {
		...(await breadcrumbEntry(parent, {
			url: route('/admin/flight-schedules/[id]', { params: { id: flightSchedule.id.toString() } }),
			title: flightTitle(flightSchedule),
		})),
		flightSchedule,
		form: await superValidate(
			existingFlightScheduleToFormData(flightSchedule),
			typebox(formSchema),
		),
	}
}
