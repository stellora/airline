import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flightInstanceID = parseInt(params.id)
	return {
		flightInstance: await apiClient
			.GET('/flight-instances/{id}', { params: { path: { id: flightInstanceID } } })
			.then((resp) => resp.data!),
		seatAssignments: await apiClient
			.GET('/flight-instances/{flightInstanceID}/seat-assignments', {
				params: { path: { flightInstanceID } },
			})
			.then((resp) => resp.data!),
	}
}
