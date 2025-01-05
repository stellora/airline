import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const route = await apiClient
		.GET('/routes/{route}', { fetch, params: { path: { route: params.route } } })
		.then((resp) => resp.data!)
	return {
		route,
	}
}
