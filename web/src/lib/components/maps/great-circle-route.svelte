<script lang="ts">
	import type { Point } from '$lib/types'
	import * as d3 from 'd3'
	import type { FeatureCollection } from 'geojson'
	import _worldMapGeoJSONData from './world-map.geojson.json'

	const worldMapGeoJSONData = _worldMapGeoJSONData as FeatureCollection

	const { origin, destination }: { origin: Point; destination: Point } = $props()

	let mapWrapper: HTMLElement | undefined

	const width = 960
	const height = 500

	const centerLat = 0
	const centerLong = -122

	// TODO!(sqs): use https://www.d3indepth.com/geographic/
	// https://connorrothschild.github.io/v4/post/svelte-and-d3
	//
	// TODO!(sqs): use topojson, more efficient https://github.com/topojson/topojson

	const geoPath = d3.geoPath(d3.geoEquirectangular().rotate([60, 0, 0]), null)

	worldMapGeoJSONData.features.push(
		{
			type: 'Feature',
			properties: null,
			geometry: {
				type: 'LineString',
				coordinates: [
					[origin.longitude, origin.latitude],
					[destination.longitude, destination.latitude]
				]
			}
		},
		{
			type: 'Feature',
			properties: null,
			geometry: {
				type: 'Point',
				coordinates: [origin.longitude, origin.latitude]
			}
		},
		{
			type: 'Feature',
			properties: null,
			geometry: {
				type: 'Point',
				coordinates: [destination.longitude, destination.latitude]
			}
		}
	)

	const svgElements = []
	for (const feature of worldMapGeoJSONData.features) {
		const p = geoPath(feature)
		if (feature.geometry.type === 'LineString') {
			svgElements.push(`<path d="${p}" fill="none" stroke="var(--map-line)" stroke-width="0.5"/>`)
		} else if (feature.geometry.type === 'Point') {
			svgElements.push(`<path d="${p}" fill="var(--map-point)" />`)
		} else if (feature.geometry.type === 'Polygon' || feature.geometry.type === 'MultiPolygon') {
			svgElements.push(
				`<path d="${p}" fill="var(--land-color)" stroke="var(--border-color)" stroke-width="0.5"/>`
			)
		}
	}

	const svgContent = `
	<svg width="${width}" height="${height}" viewBox="0 0 ${width} ${height}" xmlns="http://www.w3.org/2000/svg">
		${svgElements.join('\n')}
	</svg>
`
</script>

<div class="map-wrapper h-auto" bind:this={mapWrapper}>
	{@html svgContent}
</div>

<style>
	.map-wrapper :global(svg) {
		width: 100%;
		height: auto;

		--land-color: hsl(var(--map-land));
		--border-color: hsl(var(--map-border));

		> :global(path) {
			stroke-linejoin: bevel;
		}
	}
</style>
