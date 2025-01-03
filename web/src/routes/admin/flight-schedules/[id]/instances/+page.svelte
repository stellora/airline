<script lang="ts">
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import FlightInstanceList from '../../../flight-instances/flight-instance-list.svelte'

	let { data } = $props()
</script>

<Page title={`${flightTitle(data.flightSchedule)} instances`}>
	<FlightInstanceList flightInstances={data.flightInstances} />
</Page>
