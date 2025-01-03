<script lang="ts">
	import * as Form from '$lib/components/ui/form'
	import type { DaysOfWeek } from '$lib/types'
	import { Checkbox } from './ui/checkbox'

	let {
		value = $bindable(undefined),
	}: {
		value: DaysOfWeek | undefined
	} = $props()

	const dayNames = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
</script>

<div class="flex flex-wrap gap-3 items-center">
	{#each dayNames as dayName, day (dayName)}
		<Form.Control>
			{#snippet children({ props })}
				<Checkbox
					{...props}
					bind:checked={() => value?.includes(day as DaysOfWeek[number]),
					(v) => {
						if (v) {
							value = (value ?? []).concat(day as DaysOfWeek[number]).toSorted()
						} else {
							value = value?.filter((d) => d !== day).toSorted()
						}
					}}
					value={day.toString()}
				/>
				<Form.Label class="font-normal">
					{dayName}
				</Form.Label>
			{/snippet}
		</Form.Control>
	{/each}
</div>
