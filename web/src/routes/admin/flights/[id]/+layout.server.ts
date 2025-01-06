import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Flight, Schedule } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flights/{id}', {
		params: { path: { id } },
		fetch,
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Flight not found')
	}
	const flight = resp.data

	let schedule: Schedule | undefined
	if (flight.scheduleID) {
		const resp = await apiClient.GET('/schedules/{id}', {
			params: { path: { id: flight.scheduleID } },
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			throw error(resp.response.status, resp.error)
		}
		schedule = resp.data
	}

	return {
		flight,
		schedule,
		form: await superValidate(
			existingFlightToFormData(flight),
			typebox(schema['/flights/{id}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingFlightToFormData(
	a: Flight,
): Static<(typeof schema)['/flights/{id}']['PATCH']['args']['properties']['body']> {
	return a.scheduleID
		? {
				aircraft: a.aircraft?.registration,
				notes: a.notes,
			}
		: {
				number: a.number,
				originAirport: a.originAirport.iataCode,
				destinationAirport: a.destinationAirport.iataCode,
				fleet: a.fleet.code,
				aircraft: a.aircraft?.registration,
				departureDateTime: a.departureDateTime,
				arrivalDateTime: a.arrivalDateTime,
				notes: a.notes,
				published: a.published,
			}
}
