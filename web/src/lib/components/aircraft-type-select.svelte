<script lang="ts">
	import { page } from '$app/state'
	import * as Select from '$lib/components/ui/select/index.js'
	import { cn } from '$lib/utils'
	import type { ComponentProps } from 'svelte'
	import type { LayoutData } from '../../routes/$types'

	let {
		name,
		value = $bindable(undefined),
		class: className,
		...restProps
	}: Pick<Extract<ComponentProps<typeof Select.Root>, { type: 'single' }>, 'name' | 'value'> &
		ComponentProps<typeof Select.Trigger> = $props()

	const layoutData = page.data as unknown as LayoutData
</script>

<Select.Root type="single" {name} bind:value>
	<Select.Trigger class={cn('w-24', className)} {...restProps}>
		{#if value}
			<span class="font-mono">{value}</span>
		{/if}
	</Select.Trigger>
	<Select.Content align="start">
		{#each layoutData.allAircraftTypes as aircraftType (aircraftType.icaoCode)}
			<Select.Item value={aircraftType.icaoCode} label={aircraftType.icaoCode}
				><span class="font-mono">{aircraftType.icaoCode}</span>
				<span class="text-muted-foreground">&nbsp;&mdash; {aircraftType.name}</span></Select.Item
			>
		{/each}
	</Select.Content>
</Select.Root>
