<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { FlightSchedule } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineCode from './airline-code.svelte'
	import AirportCode from './airport-code.svelte'

	const {
		flight,
		prefix,
		showRoute = true,
		link = false,
		class: className,
		subtitleClass,
		as = 'h2',
	}: {
		flight: Pick<FlightSchedule, 'id' | 'number' | 'published'> & {
			airline: Pick<FlightSchedule['airline'], 'iataCode' | 'name'>
			originAirport: Pick<FlightSchedule['originAirport'], 'iataCode' | 'name'>
			destinationAirport: Pick<FlightSchedule['destinationAirport'], 'iataCode' | 'name'>
		}
		prefix?: string
		showRoute?: boolean
		link?: boolean
		class?: HTMLAttributes<never>['class']
		subtitleClass?: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-baseline gap-1.5')}>
	{prefix}
	<span
		class={{
			'underline decoration-dotted decoration-2 decoration-muted-foreground italic text-muted-foreground':
				!flight.published,
			'font-mono tracking-tight whitespace-nowrap': true,
		}}
	>
		{#if link}
			<a
				href={route('/admin/flight-schedules/[id]', {
					params: { id: flight.id.toString() },
				})}><AirlineCode airline={flight.airline} />&ThinSpace;{flight.number}</a
			>
		{:else}
			<AirlineCode airline={flight.airline} />&ThinSpace;{flight.number}
		{/if}
	</span>

	{#if showRoute}
		<span class={cn('text-muted-foreground text-sm whitespace-nowrap', subtitleClass)}
			><AirportCode airport={flight.originAirport} />&ndash;<AirportCode
				airport={flight.destinationAirport}
			/></span
		>
	{/if}
</svelte:element>
