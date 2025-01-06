import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flightID = parseInt(params.id)
	return {
		flight: await apiClient
			.GET('/flights/{id}', { params: { path: { id: flightID } } })
			.then((resp) => resp.data!),
		seatAssignments: await apiClient
			.GET('/flights/{flightID}/seat-assignments', {
				params: { path: { flightID } },
			})
			.then((resp) => resp.data!),
	}
}
