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

	const aircraftTypesResp = await apiClient.GET('/aircraft-types', {
		fetch,
	})
	if (!aircraftTypesResp.response.ok || !aircraftTypesResp.data) {
		throw error(aircraftTypesResp.response.status, 'Error fetching aircraft types')
	}

	return {
		allAirlines: airlinesResp.data,
		allAircraftTypes: aircraftTypesResp.data,
	}
}
