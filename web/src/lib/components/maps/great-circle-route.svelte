<script lang="ts">
	import type { Point } from '$lib/types'
	import WorldMapSVG from './web-mercator-world-map.svg?raw'

	const { origin, destination }: { origin: Point; destination: Point } = $props()
</script>

<div class="map-wrapper">
	{@html WorldMapSVG}
</div>

<style>
	.map-wrapper :global(svg) {
		--water-color: hsl(var(--background));
		--land-color: hsl(var(--primary));
		--land-border-color: hsl(var(--ring));
		--water-border-color: hsl(var(--border));
	}
</style>
