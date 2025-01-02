<script lang="ts">
	import FlightTitle from '$lib/components/flight-title.svelte'
	import type { ComponentProps } from 'svelte'
	import { fade } from 'svelte/transition'

	const {
		flightSchedule,
	}: {
		flightSchedule: ComponentProps<typeof FlightTitle>['flight']
	} = $props()
</script>

<li in:fade out:fade class="flex flex-col gap-4 border p-3 rounded-md">
	<FlightTitle flight={flightSchedule} link class="text-lg font-bold leading-none" />
</li>
