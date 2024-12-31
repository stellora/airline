<script lang="ts">
	import type { ComponentProps } from 'svelte'
	import FlightSparkRoute from './flight-spark-route.svelte'
	import FlightTitle from './flight-title.svelte'

	const {
		flight,
	}: {
		flight: ComponentProps<typeof FlightTitle>['flight'] &
			ComponentProps<typeof FlightSparkRoute>['flight']
	} = $props()
</script>

<li class="flex flex-col gap-4 border p-3 rounded-md min-h-32">
	<FlightTitle {flight} class="text-lg font-bold leading-none" />
	<FlightSparkRoute {flight} />
</li>
