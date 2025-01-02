<script lang="ts">
	import { page } from '$app/state'
	import * as Tooltip from '$lib/components/ui/tooltip'
	import type { AircraftType } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'
	import type { LayoutData } from '../../routes/$types'

	const {
		aircraftType: aircraftTypeInput,
		tooltip = true,
		class: className,
		as = 'abbr',
	}: {
		aircraftType: Pick<AircraftType, 'icaoCode' | 'name'> | AircraftType['icaoCode']
		tooltip?: boolean
		class?: HTMLAttributes<never>['class']
		as?: 'abbr'
	} = $props()

	const layoutData = page.data as unknown as LayoutData
	const aircraftType: AircraftType =
		typeof aircraftTypeInput === 'string'
			? (layoutData.allAircraftTypes.find(({ icaoCode }) => icaoCode === aircraftTypeInput) ?? {
					icaoCode: aircraftTypeInput,
					name: 'Unknown',
				})
			: aircraftTypeInput
</script>

<Tooltip.Root disabled={!tooltip}>
	<Tooltip.Trigger>
		{#snippet child({ props })}
			<svelte:element
				this={as}
				class={cn(className, 'font-mono hover:underline hover:decoration-dotted cursor-help')}
				{...props}
			>
				{aircraftType.icaoCode}
			</svelte:element>
		{/snippet}
	</Tooltip.Trigger><Tooltip.Portal>
		<Tooltip.Content collisionPadding={50}>
			<span class="font-normal text-sm">{aircraftType.name}</span>
		</Tooltip.Content>
	</Tooltip.Portal>
</Tooltip.Root>
