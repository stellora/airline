<script lang="ts">
	import type { Airport } from '$lib/types'
	import type { Feature, MultiLineString, Point } from 'geojson'
	import type { ComponentProps } from 'svelte'
	import WorldMap from './world-map.svelte'

	let {
		routes,
		...restProps
	}: { routes: [Airport, Airport][] } & Omit<
		ComponentProps<typeof WorldMap>,
		'features' | 'center' | 'fit'
	> = $props()

	let greatCircleLines: Feature<MultiLineString> = $derived({
		type: 'Feature',
		properties: null,
		geometry: {
			type: 'MultiLineString',
			coordinates: routes.map(([origin, destination]) => [
				[origin.point.longitude, origin.point.latitude],
				[destination.point.longitude, destination.point.latitude],
			]),
		},
	})

	let airports: Record<string, Airport> = $derived.by(() => {
		const airports: Record<string, Airport> = {}
		for (const [origin, destination] of routes) {
			airports[origin.iataCode] = origin
			airports[destination.iataCode] = destination
		}
		return airports
	})
</script>

<WorldMap
	features={[
		greatCircleLines,
		...Object.values(airports).map(
			(airport) =>
				({
					type: 'Feature',
					properties: { label: airport.iataCode },
					geometry: {
						type: 'Point',
						coordinates: [airport.point.longitude, airport.point.latitude],
					},
				}) satisfies Feature<Point>,
		),
	]}
	center={greatCircleLines}
	fit={greatCircleLines}
	{...restProps}
/>
