import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const schedules = apiClient
		.GET('/airports/{airportSpec}/schedules', {
			params: { path: { airportSpec: params.airportSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		schedules,
	}
}
