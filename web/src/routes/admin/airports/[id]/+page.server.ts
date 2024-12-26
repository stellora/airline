import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const airport = (
		await apiClient.GET('/airports/{id}', { params: { path: { id: params.airport } }, fetch })
	).data
	if (!airport) {
		return fail(404)
	}
	const { flightsInAirport, flightsNotInAirport } = (
		await apiClient.GET('/airports/{airportId}/flights', {
			params: { path: { airportId: params.airport } },
			fetch
		})
	).data!
	return {
		airport,
		flightsInAirport,
		flightsNotInAirport
	}
}

export const actions: Actions = {
	setFlightInAirport: async ({ request }) => {
		const data = await request.formData()

		const airport = data.get('airport')
		const flight = data.get('flight')
		if (!airport || typeof airport !== 'string') {
			return fail(400, { error: 'airport is required' })
		}
		if (!flight || typeof flight !== 'string') {
			return fail(400, { error: 'flight is required' })
		}

		const valueStr = data.get('value')
		if (valueStr !== 'true' && valueStr !== 'false') {
			return fail(400, {
				value: false,
				error: 'value must be "true" or "false"'
			})
		}
		const value = valueStr === 'true'

		const resp = await apiClient.PUT('/flights/{flightId}/airports/{airportId}', {
			params: { path: { flightId: flight, airportId: airport } },
			body: { value },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text()
			})
		}
	},
	delete: async ({ request }) => {
		const data = await request.formData()
		const id = data.get('id')
		if (!id || typeof id !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		const resp = await apiClient.DELETE('/airports/{id}', {
			params: { path: { id } },
			fetch
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text()
			})
		}
		return redirect(303, '/admin/airports')
	}
}
