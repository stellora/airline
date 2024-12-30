<script lang="ts">
	import type { Point } from '$lib/types'
	import * as d3 from 'd3'
	import type { Feature, FeatureCollection, LineString } from 'geojson'
	import _worldMapGeoJSONData from './world-map.geojson.json'

	// TODO!(sqs): inefficient
	const worldMapGeoJSONData = JSON.parse(JSON.stringify(_worldMapGeoJSONData)) as FeatureCollection

	const { origin, destination }: { origin: Point; destination: Point } = $props()

	const width = 960
	const height = 500

	// TODO!(sqs): use https://www.d3indepth.com/geographic/
	// https://connorrothschild.github.io/v4/post/svelte-and-d3
	//
	// TODO!(sqs): use topojson, more efficient https://github.com/topojson/topojson

	const greatCircleLine: Feature<LineString> = {
		type: 'Feature',
		properties: null,
		geometry: {
			type: 'LineString',
			coordinates: [
				[origin.longitude, origin.latitude],
				[destination.longitude, destination.latitude]
			]
		}
	}

	// Find midpoint longitude of the great circle line.
	const lineCentroid = d3.geoCentroid(greatCircleLine)
	const projection = d3.geoEquirectangular().rotate([lineCentroid[0], 0])
	// .fitExtent(
	// 	[
	// 		[width * 0.1, height * 0.1],
	// 		[width * 0.9, height * 0.9]
	// 	],
	// 	greatCircleLine
	// )
	const geoPath = d3.geoPath(projection, null)

	// TODO!(sqs): figure out how to rotate

	worldMapGeoJSONData.features.push(
		greatCircleLine,
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

<div class="map-wrapper h-auto">
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
