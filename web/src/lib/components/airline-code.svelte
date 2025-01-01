<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import { route } from '$lib/route-helpers'
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		airline,
		link = false,
		tooltip = true,
		class: className,
		as = 'abbr',
	}: {
		airline: Pick<Airline, 'name' | 'iataCode'>
		link?: boolean
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
				{#if link}
					<a
						href={route('/admin/airlines/[airlineSpec]', {
							params: { airlineSpec: airline.iataCode },
						})}>{airline.iataCode}</a
					>
				{:else}
					{airline.iataCode}
				{/if}
			</svelte:element>
		{/snippet}
	</Tooltip.Trigger><Tooltip.Portal>
		<Tooltip.Content collisionPadding={50}>
			<span class="font-normal text-sm">{airline.name}</span>
		</Tooltip.Content>
	</Tooltip.Portal>
</Tooltip.Root>
