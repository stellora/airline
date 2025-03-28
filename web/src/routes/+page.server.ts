import { apiClient } from '$lib/api'
import { isFeatureFlagEnabled } from '$lib/feature-flags'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ fetch }) => {
	const data: Record<string, any> = {
		flights: (await apiClient.GET('/schedules', { fetch })).data,
	}

	if (isFeatureFlagEnabled('search.nearby-airports')) {
		// Would fetch nearby airports based on geolocation
		// data.nearbyAirports = await apiClient.GET('/airports/nearby', { fetch, params: { lat, lng, radius: 100 } })
	}

	return data
}
