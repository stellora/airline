import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flightInstances = await apiClient
		.GET('/flight-schedules/{id}/instances', {
			params: { path: { id: Number.parseInt(params.id) } },
			fetch,
		})
		.then((resp) => resp.data!)
	return {
		flightInstances,
	}
}
