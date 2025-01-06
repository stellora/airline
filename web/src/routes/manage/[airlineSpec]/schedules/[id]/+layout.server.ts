import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Schedule } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flight-schedules/{id}', {
		params: { path: { id } },
		fetch,
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Schedule not found')
	}
	const schedule = resp.data
	return {
		schedule,
		form: await superValidate(
			existingScheduleToFormData(schedule),
			typebox(schema['/flight-schedules/{id}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingScheduleToFormData(
	a: Schedule,
): Static<(typeof schema)['/flight-schedules/{id}']['PATCH']['args']['properties']['body']> {
	return {
		number: a.number,
		originAirport: a.originAirport.iataCode,
		destinationAirport: a.destinationAirport.iataCode,
		fleet: a.fleet.code,
		startDate: a.startDate,
		endDate: a.endDate,
		daysOfWeek: a.daysOfWeek,
		departureTime: a.departureTime,
		durationSec: a.durationSec,
		published: a.published,
	}
}
