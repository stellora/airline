<script lang="ts">
	import type { Flight } from '$lib/types'
	import ChevronsDown from 'lucide-svelte/icons/chevrons-down'
	import ItinerarySegment from './itinerary-segment.svelte'

	const { flights }: { flights: Flight[] } = $props()
</script>

<ol class="flex flex-col">
	{#each flights as flight, i (flight.id)}
		<li>
			{#if i > 0}
				<div class="flex w-100 justify-center my-3">
					<ChevronsDown class="-ml-[205px] text-muted-foreground" />
				</div>
			{/if}
			<ItinerarySegment {flight} />
		</li>
	{/each}
</ol>
