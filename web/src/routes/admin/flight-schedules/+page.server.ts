import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flightSchedules = await apiClient
		.GET('/flight-schedules', { fetch })
		.then((resp) => resp.data!)
	return {
		flightSchedules,
	}
}
