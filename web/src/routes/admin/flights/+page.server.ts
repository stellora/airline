import { apiClient } from '$lib/api'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flights = apiClient.GET('/flights', { fetch }).then((resp) => resp.data!)
	return {
		flights: await flights
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData()

		const number = data.get('number')
		if (number === null || typeof number !== 'string') {
			return fail(400, {
				number,
				error: 'flight number is required'
			})
		}

		const originAirport = data.get('originAirport')
		const destinationAirport = data.get('destinationAirport')
		if (
			originAirport === null ||
			destinationAirport === null ||
			typeof originAirport !== 'string' ||
			typeof destinationAirport !== 'string'
		) {
			return fail(400, {
				originAirport,
				destinationAirport,
				error: 'airport code is required'
			})
		}

		const resp = await apiClient.POST('/flights', {
			body: { number, originAirport, destinationAirport },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				number,
				error: await resp.response.text()
			})
		}
	}
}
