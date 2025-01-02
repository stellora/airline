import { apiClient } from '$lib/api'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async () => {
	const airlinesResp = await apiClient.GET('/airlines', {
		fetch,
	})
	if (!airlinesResp.response.ok || !airlinesResp.data) {
		throw error(airlinesResp.response.status, 'Error fetching airlines')
	}
	return {
		allAirlines: airlinesResp.data,
	}
}
