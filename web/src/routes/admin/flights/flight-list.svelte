<script lang="ts">
	import type { Flight } from '$lib/types'
	import FlightItem from './flight-item.svelte'

	let { flights }: { flights: Flight[] } = $props()
</script>

<ul class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4">
	{#each flights as flight (flight.id)}
		<FlightItem {flight} />
	{:else}
		<p class="text-muted-foreground">No flights yet.</p>
	{/each}
</ul>
