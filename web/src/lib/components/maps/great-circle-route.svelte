<script lang="ts">
	import type { Point } from '$lib/types'
	import { project } from './generate'
	import WorldMapSVG from './world-map.svg?raw'

	const { origin, destination }: { origin: Point; destination: Point } = $props()

	let mapWrapper: HTMLDivElement

	$effect(() => {
		const svg = mapWrapper?.querySelector('svg')
		if (svg) {
			const line = document.createElementNS('http://www.w3.org/2000/svg', 'line')
			const [x1, y1] = project(origin.longitude, origin.latitude)
			const [x2, y2] = project(destination.longitude, destination.latitude)
			line.setAttribute('x1', `${x1}px`)
			line.setAttribute('y1', `${y1}px`)
			line.setAttribute('x2', `${x2}px`)
			line.setAttribute('y2', `${y2}px`)
			line.setAttribute('stroke', 'hsla(var(--primary))')
			line.setAttribute('stroke-width', '2')
			svg.appendChild(line)

			const zoom = true
			if (zoom) {
				const padding = 25
				const width = Math.abs(x2 - x1) + padding * 2
				const height = Math.abs(y2 - y1) + padding * 2
				const minX = Math.min(x1, x2) - padding
				const minY = Math.min(y1, y2) - padding
				svg.setAttribute('viewBox', `${minX} ${minY} ${width} ${height}`)
			}
		}
	})
</script>

<div class="map-wrapper h-auto" bind:this={mapWrapper}>
	{@html WorldMapSVG}
</div>

<style>
	.map-wrapper :global(svg) {
		width: 100%;
		height: auto;

		--land-color: hsl(var(--foreground) / 50%);
		--border-color: hsl(var(--muted-foreground));

		> :global(path) {
			stroke-linejoin: bevel;
		}
	}
</style>
