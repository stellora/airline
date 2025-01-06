<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import { route } from '$lib/route-helpers'
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineIcon from './airline-icon.svelte'

	const {
		airline,
		link = false,
		tooltip = true,
		showName = false,
		icon = false,
		class: className,
		as = 'abbr',
	}: {
		airline: Pick<Airline, 'name' | 'iataCode'>
		link?: boolean
		tooltip?: boolean
		showName?: boolean
		icon?: boolean
		class?: HTMLAttributes<never>['class']
		as?: 'abbr'
	} = $props()
</script>

<Tooltip.Root disabled={!tooltip}>
	<Tooltip.Trigger>
		{#snippet child({ props })}
			<svelte:element
				this={as}
				class={cn(className, 'font-mono inline-flex items-baseline gap-1.5 overflow-hidden', {
					'hover:underline hover:decoration-dotted': tooltip,
					'cursor-help': !link && tooltip,
				})}
				{...props}
			>
				{#if icon}<AirlineIcon {airline} class="self-center -mt-0.5 flex-shrink-0" />{/if}
				{#if link}
					<a
						href={route('/admin/airlines/[airlineSpec]', {
							params: { airlineSpec: airline.iataCode },
						})}>{airline.iataCode}</a
					>
				{:else}
					{airline.iataCode}
				{/if}
				{#if showName}
					<span data-airline-name class="text-sm font-sans text-muted-foreground truncate"
						>{airline.name}</span
					>
				{/if}
			</svelte:element>
		{/snippet}
	</Tooltip.Trigger><Tooltip.Portal>
		<Tooltip.Content collisionPadding={50}>
			<span class="font-normal">{airline.name}</span>
		</Tooltip.Content>
	</Tooltip.Portal>
</Tooltip.Root>
