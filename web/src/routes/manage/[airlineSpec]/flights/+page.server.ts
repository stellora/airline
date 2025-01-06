import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	return {
		flightInstances: await apiClient
			.GET('/airlines/{airlineSpec}/flight-instances', {
				params: { path: { airlineSpec: params.airlineSpec } },
				fetch,
			})
			.then((resp) => resp.data),
	}
}
