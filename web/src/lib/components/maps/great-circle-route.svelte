<script lang="ts">
	import type { Point } from '$lib/types'
	import * as d3 from 'd3'
	import type { Feature, FeatureCollection, LineString } from 'geojson'
	import _worldMapGeoJSONData from './world-map.geojson.json'

	// TODO!(sqs): inefficient
	const worldMapGeoJSONData = JSON.parse(JSON.stringify(_worldMapGeoJSONData)) as FeatureCollection

	const { origin, destination }: { origin: Point; destination: Point } = $props()

	let containerRef: HTMLDivElement | undefined
	let width = $state(960)

	$effect(() => {
		if (!containerRef) return
		const resizeObserver = new ResizeObserver((entries) => {
			width = entries[0].contentRect.width
		})
		resizeObserver.observe(containerRef)
		return () => resizeObserver.disconnect()
	})
	// TODO!(sqs): use https://www.d3indepth.com/geographic/
	// https://connorrothschild.github.io/v4/post/svelte-and-d3
	//
	// TODO!(sqs): use topojson, more efficient https://github.com/topojson/topojson

	function makeSVG(): string {
		const height = width / 1.92

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

		const lineCentroid = d3.geoCentroid(greatCircleLine)
		const padding = 50
		const projection = d3
			.geoEquirectangular()
			.rotate([-1 * lineCentroid[0], 0])
			.fitExtent(
				[
					[padding, padding],
					[width - padding, height - padding]
				],
				greatCircleLine
			)
		const geoPath = d3.geoPath(projection, null)

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
	<svg viewBox="0 0 ${width} ${height}" xmlns="http://www.w3.org/2000/svg">
		${svgElements.join('\n')}
	</svg>`
		return svgContent
	}
	const svgContent = $derived(makeSVG())
</script>

<div class="map-wrapper h-auto" bind:this={containerRef}>
	{@html svgContent}
</div>

<style>
	.map-wrapper {
		overflow: hidden;
		display: flex;
	}
	.map-wrapper :global(svg) {
		height: auto;

		--land-color: hsl(var(--map-land));
		--border-color: hsl(var(--map-border));
	}
</style>
