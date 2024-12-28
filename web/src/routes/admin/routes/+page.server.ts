import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const routes = apiClient.GET('/routes', { fetch }).then((resp) => resp.data!)
	return {
		routes
	}
}
