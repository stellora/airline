<script lang="ts">
	import { flightRoute } from '$lib/flight-helpers.js'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-4xl font-bold">
		{flightRoute(data.route.originAirport, data.route.destinationAirport)}
	</h1>
</div>
