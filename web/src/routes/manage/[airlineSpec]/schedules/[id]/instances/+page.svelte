<script lang="ts">
	import FlightTitle from '$lib/components/flight-title.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import FlightInstanceList from '../../../../../admin/flight-instances/flight-instance-list.svelte'

	let { data } = $props()
</script>

<Page title={`${flightTitle(data.schedule)} instances`}>
	<FlightTitle flight={data.schedule} as="span" showRoute={true} />
	<FlightInstanceList flightInstances={data.flightInstances} />
</Page>
