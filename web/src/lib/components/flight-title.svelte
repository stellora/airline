<script lang="ts">
	import type { Flight } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		flight,
		link = false,
		class: className,
		as = 'h2'
	}: {
		flight: Pick<Flight, 'id' | 'number' | 'published'> & {
			originAirport: Pick<Flight['originAirport'], 'iataCode'>
			destinationAirport: Pick<Flight['destinationAirport'], 'iataCode'>
		}
		link?: boolean
		class?: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-baseline gap-1.5')}>
	<span
		class={{
			'underline decoration-dotted decoration-2 decoration-muted-foreground italic text-muted-foreground':
				!flight.published
		}}
	>
		{#if link}
			<a href={`/admin/flights/${flight.id}`}>{flight.number}</a>
		{:else}
			{flight.number}
		{/if}
	</span>

	<span class="text-muted-foreground text-sm whitespace-nowrap"
		>{flight.originAirport.iataCode}&ndash;{flight.destinationAirport.iataCode}</span
	>
</svelte:element>
