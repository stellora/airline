import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flights = await apiClient
		.GET('/schedules/{id}/flights', {
			params: { path: { id: Number.parseInt(params.id) } },
			fetch,
		})
		.then((resp) => resp.data!)
	return {
		flights,
	}
}
