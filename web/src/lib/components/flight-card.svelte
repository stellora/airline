<script lang="ts">
	import type { Airport, Flight } from '$lib/types'
	import Tag from 'lucide-svelte/icons/tag'
	import FlightTitle from './flight-title.svelte'

	const { flight, airports }: { flight: Flight; airports?: Airport[] } = $props()
</script>

<li class="flex flex-col gap-4 border p-3 rounded-md min-h-32">
	<FlightTitle {flight} class="text-lg font-bold leading-none" />
	{#if airports && airports.length > 0}
		<div class="flex gap-2 items-start">
			<Tag size="1.5ch" class="mt-1.5 text-muted-foreground" />
			<ul class="flex gap-2">
				{#each airports as airport (airport.id)}
					<span class="p-1 text-xs rounded border">{airport.title}</span>
				{/each}
			</ul>
		</div>
	{/if}
</li>
