<script lang="ts">
	import type { Point } from '$lib/types'
	import { debounce } from '$lib/utils'
	import * as d3 from 'd3'
	import type { Feature, FeatureCollection, LineString } from 'geojson'
	import { feature as topojsonToGeoJSON } from 'topojson-client'
	import worldTopoJSONData from './world.topojson.json'

	const {
		origin,
		originLabel,
		destination,
		destinationLabel
	}: { origin: Point; originLabel?: string; destination: Point; destinationLabel?: string } =
		$props()

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
	const featureCollections: FeatureCollection[] = [
		topojsonToGeoJSON(
			worldTopoJSONData as unknown as TopoJSON.Topology,
			'world'
		) as unknown as FeatureCollection,
		{
			type: 'FeatureCollection',
			features: [
				greatCircleLine,
				{
					type: 'Feature',
					properties: { label: originLabel },
					geometry: {
						type: 'Point',
						coordinates: [origin.longitude, origin.latitude]
					}
				},
				{
					type: 'Feature',
					properties: { label: destinationLabel },
					geometry: {
						type: 'Point',
						coordinates: [destination.longitude, destination.latitude]
					}
				}
			]
		}
	]
	const lineCentroid = d3.geoCentroid(greatCircleLine)

	// Dynamically scale SVG.
	let containerRef: HTMLDivElement | undefined
	let width = $state(960)
	let height = $derived(width / 1.92)
	$effect(() => {
		if (!containerRef) return
		const resizeObserver = new ResizeObserver(
			debounce((entries) => {
				width = entries[0].contentRect.width
			}, 25)
		)
		resizeObserver.observe(containerRef)
		return () => resizeObserver.disconnect()
	})

	// TODO!(sqs): use https://www.d3indepth.com/geographic/
	// https://connorrothschild.github.io/v4/post/svelte-and-d3
	//
	// TODO!(sqs): use topojson, more efficient https://github.com/topojson/topojson

	const geoPath = $derived.by(() => {
		const padding = [0.08 * width, 0.08 * height]
		const projection = d3
			.geoEquirectangular()
			.rotate([-1 * lineCentroid[0], 0])
			.fitExtent(
				[
					[padding[0], padding[1]],
					[width - padding[0], height - padding[1]]
				],
				greatCircleLine
			)
			.clipExtent([
				[0, 0],
				[width, height]
			])
		return d3.geoPath(projection, null)
	})
</script>

<div class="map-wrapper h-auto" bind:this={containerRef}>
	<svg viewBox="0 0 {width} {height}" xmlns="http://www.w3.org/2000/svg">
		{#each featureCollections as collection}
			{#each collection.features as feature}
				{@const path = geoPath(feature)}
				{#if path === null}{:else if feature.geometry.type === 'LineString'}
					<path d={path} fill="none" stroke="var(--map-line)" stroke-width="2.5" />
				{:else if feature.geometry.type === 'Point'}
					{#if feature.properties?.label}
						{@const [x, y] = geoPath.centroid(feature)}
						{@const charWidth = 7.5}
						{@const charHeight = 14.5}
						{@const size = [feature.properties?.label.length * charWidth, charHeight]}
						{@const padding = [3, 3]}
						<g class="font-mono text-xs leading-none">
							<rect
								x={x - size[0] / 2 - padding[0]}
								y={y - size[1] / 2 - padding[1]}
								width={size[0] + 2 * padding[0]}
								height={size[1] + 2 * padding[1]}
								fill="var(--map-point)"
								stroke-width="1"
								rx="2"
							/>
							<text {x} y={y + size[1] * 0.25} text-anchor="middle" fill="var(--map-point-text)">
								{feature.properties?.label}
							</text>
						</g>
					{:else}
						<path d={path} fill="var(--map-point)" />
					{/if}
				{:else if feature.geometry.type === 'Polygon' || feature.geometry.type === 'MultiPolygon'}
					<path d={path} fill="var(--land-color)" stroke="var(--border-color)" stroke-width="0.5" />
				{/if}
			{/each}
		{/each}
	</svg>
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
