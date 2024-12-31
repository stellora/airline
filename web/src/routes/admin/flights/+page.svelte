<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'

	let { data } = $props()
</script>

<Page title="Flights">
	{#snippet actions()}
		<Button href="/admin/flights/new" variant="default"><Plus /> New flight</Button>
	{/snippet}
	<FlightList flights={data.flights} />
</Page>
