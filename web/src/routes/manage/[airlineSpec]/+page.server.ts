import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const schedules = apiClient
		.GET('/airlines/{airlineSpec}/schedules', {
			params: { path: { airlineSpec: params.airlineSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		schedules,
	}
}
