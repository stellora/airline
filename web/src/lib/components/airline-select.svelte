<script lang="ts">
	import { page } from '$app/state'
	import * as Select from '$lib/components/ui/select/index.js'
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { ComponentProps } from 'svelte'
	import type { LayoutData } from '../../routes/$types'
	import AirlineCode from './airline-code.svelte'

	let {
		name,
		value = $bindable(undefined),
		showName = true,
		class: className,
		...restProps
	}: Pick<Extract<ComponentProps<typeof Select.Root>, { type: 'single' }>, 'name'> & {
		value: Airline | undefined
		showName?: boolean
	} & Omit<ComponentProps<typeof Select.Trigger>, 'value'> = $props()

	const layoutData = page.data as unknown as LayoutData
</script>

<Select.Root
	type="single"
	{name}
	bind:value={() => value?.iataCode ?? '',
	(iataCode) => {
		value = layoutData.allAirlines.find((airline) => airline.iataCode === iataCode)
	}}
	allowDeselect={false}
>
	<Select.Trigger
		class={cn('w-auto', showName ? 'min-w-[250px]' : 'min-w-[90px]', className)}
		{...restProps}
	>
		<span class="pr-1.5"
			>{#if value}
				<AirlineCode airline={value} tooltip={false} icon {showName} />
			{/if}</span
		>
	</Select.Trigger>
	<Select.Content align="start">
		{#each layoutData.allAirlines as airline (airline.id)}
			<Select.Item value={airline.iataCode} label={airline.iataCode}
				><AirlineCode {airline} tooltip={false} icon {showName} />
			</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
