import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ fetch }) => {
	return {
		flights: (await apiClient.GET('/schedules', { fetch })).data,
	}
}
