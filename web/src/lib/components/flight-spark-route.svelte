<script lang="ts">
	import type { Flight } from '$lib/types'
	import type { Feature, LineString } from 'geojson'
	import WorldMap from './maps/world-map.svelte'

	const { flight }: { flight: Pick<Flight, 'originAirport' | 'destinationAirport'> } = $props()

	const line: Feature<LineString> = {
		type: 'Feature',
		properties: null,
		geometry: {
			type: 'LineString',
			coordinates: [
				[flight.originAirport.point.longitude, flight.originAirport.point.latitude],
				[flight.destinationAirport.point.longitude, flight.destinationAirport.point.latitude],
			],
		},
	}
</script>

<WorldMap features={[line]} fit={line} detailLevel="low" drawBorders={false} />
