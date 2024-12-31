import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flights = apiClient.GET('/flights', { fetch }).then((resp) => resp.data!)
	return {
		flights: await flights,
	}
}
