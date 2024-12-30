<script lang="ts">
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Card } from '$lib/components/ui/card'
	import { flightRoute } from '$lib/flight-helpers.js'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-4xl font-bold">
		{flightRoute(data.route.originAirport, data.route.destinationAirport)}
	</h1>
	<Card class="max-w-[960px]">
		<GreatCircleRoute routes={[[data.route.originAirport, data.route.destinationAirport]]} />
	</Card>
</div>
