import { apiClient } from '$lib/api'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/itineraries/{itinerarySpec}', {
		params: { path: { itinerarySpec: params.itinerarySpec } },
		fetch,
	})
	const itinerary = resp.data
	if (!itinerary) {
		error(resp.response.status, resp.error)
	}
	return {
		itinerary,
	}
}
