<script lang="ts">
	import type { Flight } from '$lib/types'
	import ItinerarySegment from './itinerary-segment.svelte'

	const { flights }: { flights: Flight[] } = $props()
</script>

{#each flights as flight (flight.id)}
	<ItinerarySegment {flight} />
{/each}
