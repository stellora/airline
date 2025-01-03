<script lang="ts">
	import Page from '$lib/components/ui/page/page.svelte'
	import FlightInstanceList from './flight-instance-list.svelte'

	let { data } = $props()
</script>

<Page title="Flight instances">
	<FlightInstanceList flightInstances={data.flightInstances} />
</Page>
