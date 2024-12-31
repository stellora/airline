<script lang="ts">
	import Page from '$lib/components/ui/page/page.svelte'
	import FlightList from './flight-list.svelte'

	let { data } = $props()
</script>

<Page title="Flights">
	<a href="/admin/flights/new">Add flight TODO!(sqs)</a>
	<FlightList flights={data.flights} />
</Page>
