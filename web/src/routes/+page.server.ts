import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ fetch }) => {
	return {
		products: (await apiClient.GET('/products', { fetch })).data
	}
}
