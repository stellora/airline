import { apiClient } from '$lib/api'
import type { FlightSchedule } from '$lib/types'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'
import {
	existingFlightInstanceToFormData,
	flightInstanceFromScheduleFormSchema,
} from './flight-instance-form'

export const load: LayoutServerLoad = async ({ params }) => {
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

	let flightSchedule: FlightSchedule | undefined
	if (flightInstance.scheduleID) {
		const resp = await apiClient.GET('/flight-schedules/{id}', {
			params: { path: { id: flightInstance.scheduleID } },
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			throw error(resp.response.status, resp.error)
		}
		flightSchedule = resp.data
	}

	return {
		flightInstance,
		flightSchedule,
		form: await superValidate(
			existingFlightInstanceToFormData(flightInstance),
			typebox(flightInstanceFromScheduleFormSchema),
		),
	}
}
