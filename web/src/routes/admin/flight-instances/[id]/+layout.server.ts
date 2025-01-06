import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { FlightInstance, FlightSchedule } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { LayoutServerLoad } from './$types'

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
			typebox(schema['/flight-instances/{id}']['PATCH']['args']['properties']['body']),
		),
	}
}

function existingFlightInstanceToFormData(
	a: FlightInstance,
): Static<(typeof schema)['/flight-instances/{id}']['PATCH']['args']['properties']['body']> {
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
