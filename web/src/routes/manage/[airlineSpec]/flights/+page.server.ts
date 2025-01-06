import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	return {
		flights: await apiClient
			.GET('/airlines/{airlineSpec}/flights', {
				params: { path: { airlineSpec: params.airlineSpec } },
				fetch,
			})
			.then((resp) => resp.data),
	}
}
