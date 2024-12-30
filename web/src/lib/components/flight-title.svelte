<script lang="ts">
	import type { Flight } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirportCode from './airport-code.svelte'

	const {
		flight,
		link = false,
		class: className,
		subtitleClass,
		as = 'h2',
	}: {
		flight: Pick<Flight, 'id' | 'number' | 'published'> & {
			originAirport: Exclude<Flight['originAirport'], 'name' | 'iataCode'>
			destinationAirport: Pick<Flight['destinationAirport'], 'name' | 'iataCode'>
		}
		link?: boolean
		class?: HTMLAttributes<never>['class']
		subtitleClass?: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-baseline gap-1.5')}>
	<span
		class={{
			'underline decoration-dotted decoration-2 decoration-muted-foreground italic text-muted-foreground':
				!flight.published,
			'font-mono': true,
		}}
	>
		{#if link}
			<a href={`/admin/flights/${flight.id}`}>{flight.number}</a>
		{:else}
			{flight.number}
		{/if}
	</span>

	<span class={cn('text-muted-foreground text-sm whitespace-nowrap', subtitleClass)}
		><AirportCode airport={flight.originAirport} />&ndash;<AirportCode
			airport={flight.destinationAirport}
		/></span
	>
</svelte:element>
