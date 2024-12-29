<script lang="ts">
	import type { Point } from '$lib/types'
	import WorldMapSVG from './world-map.svg?raw'

	const { origin, destination }: { origin: Point; destination: Point } = $props()
</script>

<div class="map-wrapper">
	{@html WorldMapSVG}
</div>

<style>
	.map-wrapper :global(svg) {
		--land-color: hsl(var(--foreground) / 50%);
		--border-color: hsla(var(--border) / 50%);
	}
</style>
