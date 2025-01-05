<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { FlightSchedule } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineCode from './airline-code.svelte'
	import AirportCode from './airport-code.svelte'

	let {
		flight,
		prefix,
		showAirlineIcon = true,
		showRoute = true,
		link = false,
		tooltip = true,
		class: className,
		linkClass,
		subtitleClass,
		as = 'h2',
	}: {
		flight: Pick<FlightSchedule, 'id' | 'number' | 'published'> & {
			airline: Pick<FlightSchedule['airline'], 'iataCode' | 'name'>
			originAirport: Pick<FlightSchedule['originAirport'], 'iataCode' | 'name'>
			destinationAirport: Pick<FlightSchedule['destinationAirport'], 'iataCode' | 'name'>
		}
		prefix?: string
		showAirlineIcon?: boolean
		showRoute?: boolean
		link?: boolean
		tooltip?: boolean
		class?: HTMLAttributes<never>['class']
		linkClass?: HTMLAttributes<never>['class']
		subtitleClass?: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-baseline gap-1.5')}>
	{prefix}
	<span
		class={{
			'italic text-muted-foreground': !flight.published,
			'font-mono tracking-tight whitespace-nowrap leading-none': true,
		}}
	>
		{#if link}
			<a
				href={route('/admin/flight-schedules/[id]', {
					params: { id: flight.id.toString() },
				})}
				class={linkClass}
				><AirlineCode airline={flight.airline} icon={showAirlineIcon} {tooltip} /><span
					class="font-sans text-xs"
					>&nbsp;
				</span>{flight.number}</a
			>
		{:else}
			<AirlineCode airline={flight.airline} icon={showAirlineIcon} {tooltip} /><span
				class="font-sans text-xs"
			>
				&nbsp;
			</span>{flight.number}
		{/if}
	</span>

	{#if showRoute}
		<span class={cn('text-sm whitespace-nowrap leading-none', subtitleClass)}
			><AirportCode airport={flight.originAirport} {tooltip} />&ndash;<AirportCode
				airport={flight.destinationAirport}
				{tooltip}
			/></span
		>
	{/if}
</svelte:element>
