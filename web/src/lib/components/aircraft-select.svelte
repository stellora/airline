<script lang="ts">
	import { apiClient } from '$lib/api'
	import * as Select from '$lib/components/ui/select/index.js'
	import type { AirlineSpec } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { ComponentProps } from 'svelte'

	let {
		name,
		value = $bindable(undefined),
		byAirline,
		class: className,
		...restProps
	}: Pick<Extract<ComponentProps<typeof Select.Root>, { type: 'single' }>, 'name' | 'value'> &
		ComponentProps<typeof Select.Trigger> & {
			byAirline: AirlineSpec
		} = $props()

	const aircraft = apiClient
		.GET('/airlines/{airlineSpec}/aircraft', {
			params: { path: { airlineSpec: byAirline } },
			fetch,
		})
		.then((resp) => resp.data!)
</script>

<Select.Root type="single" {name} bind:value allowDeselect={false}>
	<Select.Trigger class={cn('w-[170px]', className)} {...restProps}>
		<span class="font-mono pr-1.5"
			>{#if value}{value}{/if}</span
		>
	</Select.Trigger>
	<Select.Content align="start">
		{#await aircraft then allAircraft}
			{#each allAircraft as aircraft (aircraft.id)}
				<Select.Item value={aircraft.registration} label={aircraft.registration}
					><span class="font-mono">{aircraft.registration}</span>
					<span class="text-muted-foreground">&nbsp;&mdash; {aircraft.aircraftType}</span
					></Select.Item
				>
			{/each}
		{/await}
	</Select.Content>
</Select.Root>
