<script lang="ts">
	import type { Airport } from '$lib/types'
	import type { Feature, LineString } from 'geojson'
	import WorldMap from './world-map.svelte'

	const { origin, destination }: { origin: Airport; destination: Airport } = $props()

	const greatCircleLine: Feature<LineString> = {
		type: 'Feature',
		properties: null,
		geometry: {
			type: 'LineString',
			coordinates: [
				[origin.point.longitude, origin.point.latitude],
				[destination.point.longitude, destination.point.latitude]
			]
		}
	}
</script>

<WorldMap
	features={[
		greatCircleLine,
		{
			type: 'Feature',
			properties: { label: origin.iataCode },
			geometry: {
				type: 'Point',
				coordinates: [origin.point.longitude, origin.point.latitude]
			}
		},
		{
			type: 'Feature',
			properties: { label: destination.iataCode },
			geometry: {
				type: 'Point',
				coordinates: [destination.point.longitude, destination.point.latitude]
			}
		}
	]}
	center={greatCircleLine}
	fit={greatCircleLine}
/>
