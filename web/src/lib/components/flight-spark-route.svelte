<script lang="ts">
	import { geoDistanceMiles } from '$lib/flight-helpers'
	import type { FlightInstance, Schedule } from '$lib/types'
	import type { Feature, LineString } from 'geojson'
	import type { ComponentProps } from 'svelte'
	import WorldMap from './maps/world-map.svelte'

	const {
		flight,
		drawBorders: drawBordersArg = 'auto',
		...restProps
	}: {
		flight: Pick<FlightInstance | Schedule, 'originAirport' | 'destinationAirport'>
	} & Omit<
		ComponentProps<typeof WorldMap>,
		'features' | 'center' | 'fit' | 'detailLevel' | 'drawBorders'
	> & {
			drawBorders?: boolean | 'auto'
		} = $props()

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

	const drawBorders = drawBordersArg === 'auto' ? geoDistanceMiles(line) < 3000 : drawBordersArg
</script>

<WorldMap
	features={[line]}
	center={line}
	fit={line}
	detailLevel="auto"
	{drawBorders}
	{...restProps}
/>
