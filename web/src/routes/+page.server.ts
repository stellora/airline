import { apiClient } from '$lib/api'
import { isFeatureFlagEnabled } from '$lib/feature-flags'
import type { PageServerLoad } from './$types'

// Helper function to determine if a flight is related to Australia
function isAustralianFlight(flight: any) {
	const originCountry = flight.originAirport?.country;
	const destinationCountry = flight.destinationAirport?.country;
	return originCountry === 'Australia' || destinationCountry === 'Australia';
}

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await apiClient.GET('/schedules', { fetch });
	const flights = response.data;

	// Custom sort: Australian flights first, then others
	const sortedFlights = [...flights].sort((a, b) => {
		const aIsAustralian = isAustralianFlight(a);
		const bIsAustralian = isAustralianFlight(b);
		
		// Australian flights come first
		if (aIsAustralian && !bIsAustralian) return -1;
		if (!aIsAustralian && bIsAustralian) return 1;
		
		// Otherwise, maintain original order
		return 0;
	});

	const data: Record<string, any> = {
		flights: sortedFlights,
	}

	if (isFeatureFlagEnabled('search.nearby-airports')) {
		// Would fetch nearby airports based on geolocation
		// data.nearbyAirports = await apiClient.GET('/airports/nearby', { fetch, params: { lat, lng, radius: 100 } })
	}

	return data
}
