<script lang="ts">
	import FlightList from './flight-list.svelte'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Flights</h1>
	<a href="/admin/flights/new">Add flight TODO!(sqs)</a>
	<FlightList flights={data.flights} />
</div>
