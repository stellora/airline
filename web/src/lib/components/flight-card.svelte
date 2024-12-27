<script lang="ts">
	import type { Airport } from '$lib/types'
	import type { ComponentProps } from 'svelte'
	import FlightTitle from './flight-title.svelte'

	const {
		flight,
		airports
	}: { flight: ComponentProps<typeof FlightTitle>['flight']; airports?: Airport[] } = $props()
</script>

<li class="flex flex-col gap-4 border p-3 rounded-md min-h-32">
	<FlightTitle {flight} class="text-lg font-bold leading-none" />
</li>
