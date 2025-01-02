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
	<Select.Trigger class={cn('w-auto', className)} {...restProps}>
		{#if value}
			<span class="font-mono">{value}</span>
		{:else}
			<span class="font-mono">&nbsp;&nbsp;</span>
		{/if}
	</Select.Trigger>
	<Select.Content align="start">
		{#each layoutData.allAirlines as airline (airline.id)}
			<Select.Item value={airline.iataCode} label={airline.iataCode}
				><span class="font-mono">{airline.iataCode}</span>
				<span class="text-muted-foreground">&nbsp;&mdash; {airline.name}</span></Select.Item
			>
		{/each}
	</Select.Content>
</Select.Root>
