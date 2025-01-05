import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const fleets = await apiClient
		.GET('/airlines/{airlineSpec}/fleets', {
			params: { path: { airlineSpec: params.airlineSpec } },
			fetch,
		})
		.then((resp) => resp.data!)
	return {
		fleets,
	}
}
