import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flights = await apiClient.GET('/flights', { fetch }).then((resp) => {
		console.log(resp.error)
		return resp.data!
	})
	return {
		flights,
	}
}
