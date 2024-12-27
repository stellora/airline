<script lang="ts">
	import FlightForm from './flight-form.svelte'
	import FlightList from './flight-list.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Flights</h1>
	<FlightForm {form} />
	<FlightList flights={data.flights} />
	{data.asdf}
</div>
