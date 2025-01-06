import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { FlightSchedule } from '$lib/types'
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
		throw error(404, 'Flight schedule not found')
	}
	const flightSchedule = resp.data
	return {
		flightSchedule,
		form: await superValidate(
			existingFlightScheduleToFormData(flightSchedule),
			typebox(schema['/flight-schedules/{id}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingFlightScheduleToFormData(
	a: FlightSchedule,
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
