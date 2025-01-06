<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { Schedule } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineCode from './airline-code.svelte'
	import AirportCode from './airport-code.svelte'

	let {
		schedule,
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
		schedule: Pick<Schedule, 'id' | 'number' | 'published'> & {
			airline: Pick<Schedule['airline'], 'iataCode' | 'name'>
			originAirport: Pick<Schedule['originAirport'], 'iataCode' | 'name'>
			destinationAirport: Pick<Schedule['destinationAirport'], 'iataCode' | 'name'>
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
			'italic text-muted-foreground': !schedule.published,
			'font-mono tracking-tight whitespace-nowrap leading-none': true,
		}}
	>
		{#if link}
			<a
				href={route('/manage/[airlineSpec]/schedules/[id]', {
					params: { airlineSpec: schedule.airline.iataCode, id: schedule.id.toString() },
				})}
				class={linkClass}
				><AirlineCode airline={schedule.airline} icon={showAirlineIcon} {tooltip} /><span
					class="font-sans text-xs"
					>&nbsp;
				</span>{schedule.number}</a
			>
		{:else}
			<AirlineCode airline={schedule.airline} icon={showAirlineIcon} {tooltip} /><span
				class="font-sans text-xs"
			>
				&nbsp;
			</span>{schedule.number}
		{/if}
	</span>

	{#if showRoute}
		<span class={cn('text-sm whitespace-nowrap leading-none', subtitleClass)}
			><AirportCode airport={schedule.originAirport} {tooltip} />&ndash;<AirportCode
				airport={schedule.destinationAirport}
				{tooltip}
			/></span
		>
	{/if}
</svelte:element>
