<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import { route } from '$lib/route-helpers'
	import type { Airport } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		airport,
		link = false,
		tooltip = true,
		class: className,
		as = 'abbr',
	}: (
		| {
				airport: Pick<Airport, 'name' | 'iataCode'>
				link?: false
		  }
		| {
				airport: Pick<Airport, 'id' | 'name' | 'iataCode'>
				link: true
		  }
	) & {
		tooltip?: boolean
		class?: HTMLAttributes<never>['class']
		as?: 'abbr'
	} = $props()
</script>

<Tooltip.Root disabled={!tooltip}>
	<Tooltip.Trigger>
		{#snippet child({ props })}
			<svelte:element
				this={as}
				class={cn(className, 'font-mono hover:underline hover:decoration-dotted', {
					'cursor-help': !link,
				})}
				{...props}
			>
				{#if link && 'id' in airport}
					<a
						href={route('/admin/airports/[airportSpec]', {
							params: { airportSpec: airport.iataCode },
						})}>{airport.iataCode}</a
					>
				{:else}
					{airport.iataCode}
				{/if}
			</svelte:element>
		{/snippet}
	</Tooltip.Trigger><Tooltip.Portal>
		<Tooltip.Content collisionPadding={50}>
			<span class="font-normal text-sm">{airport.name}</span>
		</Tooltip.Content>
	</Tooltip.Portal>
</Tooltip.Root>
