<script lang="ts">
	import { debounce } from '$lib/utils'
	import * as d3 from 'd3'
	import type { Feature, FeatureCollection } from 'geojson'
	import { feature as topojsonToGeoJSON } from 'topojson-client'
	import worldTopoJSONData from './world.topojson.json'

	const {
		features,
		center: centerFeature,
		fit: fitFeature,
		width: widthArg = 'auto',
	}: { features: Feature[]; center?: Feature; fit?: Feature; width?: number | 'auto' } = $props()

	// Dynamically scale SVG.
	let containerRef: HTMLDivElement | undefined
	let width = $state(widthArg === 'auto' ? 960 : widthArg)
	let height = $derived(width / 1.92)

	$effect(() => {
		if (widthArg === 'auto') {
			if (!containerRef) return
			const resizeObserver = new ResizeObserver(
				debounce<ResizeObserverCallback>((entries) => {
					width = entries[0].contentRect.width
				}, 25),
			)
			resizeObserver.observe(containerRef)
			return () => resizeObserver.disconnect()
		}
	})

	const featureCollections: FeatureCollection[] = [
		topojsonToGeoJSON(
			worldTopoJSONData as unknown as TopoJSON.Topology,
			'world',
		) as unknown as FeatureCollection,
		{
			type: 'FeatureCollection',
			features,
		},
	]

	const centerCentroid = centerFeature && d3.geoCentroid(centerFeature)

	// TODO!(sqs): use https://www.d3indepth.com/geographic/
	// https://connorrothschild.github.io/v4/post/svelte-and-d3
	//
	// TODO!(sqs): use topojson, more efficient https://github.com/topojson/topojson

	const geoPath = $derived.by(() => {
		const padding = [0.08 * width, 0.08 * height]
		const projection = d3.geoEquirectangular()
		if (centerCentroid) {
			projection.rotate([-1 * centerCentroid[0], 0])
		}
		if (fitFeature) {
			projection
				.fitExtent(
					[
						[padding[0], padding[1]],
						[width - padding[0], height - padding[1]],
					],
					fitFeature,
				)
				.clipExtent([
					[0, 0],
					[width, height],
				])
		}
		return d3.geoPath(projection, null)
	})
</script>

<div class="map-wrapper" bind:this={containerRef}>
	<svg {width} {height} viewBox="0 0 {width} {height}" xmlns="http://www.w3.org/2000/svg">
		{#each featureCollections as collection}
			{#each collection.features as feature}
				{@const path = geoPath(feature)}
				{#if path === null}{:else if feature.geometry.type === 'LineString' || feature.geometry.type === 'MultiLineString'}
					<path d={path} fill="none" stroke="var(--map-line)" stroke-width="2.5" />
				{:else if feature.geometry.type === 'Point'}
					{#if feature.properties?.label}
						{@const [x, y] = geoPath.centroid(feature)}
						{@const charWidth = 9.5}
						<!-- 7 -->
						{@const charHeight = 11}
						<!-- 12 -->
						{@const size = [feature.properties?.label.length * charWidth, charHeight]}
						{@const padding = [5, 4]}
						<g class="font-mono text-base leading-none hover:font-bold">
							<rect
								x={x - 0.5 * size[0] - padding[0]}
								y={y - 0.5 * size[1] - padding[1]}
								width={size[0] + 2 * padding[0]}
								height={size[1] + 2 * padding[1]}
								fill="var(--map-point)"
								rx="2"
							/>
							<text {x} y={y + 0.5 * size[1]} text-anchor="middle" fill="var(--map-point-text)">
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
