<script lang="ts">
	import Page from '$lib/components/ui/page/page.svelte'
	import FlightInstanceList from '../../../admin/flight-instances/flight-instance-list.svelte'

	let { data } = $props()
</script>

<Page title={`Flights - ${data.airline.iataCode}`}>
	<FlightInstanceList flightInstances={data.flightInstances} showFlightInfo />
</Page>
