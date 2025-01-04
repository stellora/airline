import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flightInstances = await apiClient.GET('/flight-instances', { fetch }).then((resp) => {
		console.log(resp.error)
		return resp.data!
	})
	return {
		flightInstances,
	}
}
