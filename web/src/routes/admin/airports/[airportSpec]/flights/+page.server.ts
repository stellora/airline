import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flightSchedules = apiClient
		.GET('/airports/{airportSpec}/flight-schedules', {
			params: { path: { airportSpec: params.airportSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		flightSchedules,
	}
}
