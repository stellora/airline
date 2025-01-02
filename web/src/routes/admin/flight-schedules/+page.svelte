<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'

	let { data } = $props()
</script>

<Page title="Flight schedules">
	{#snippet actions()}
		<Button href={route('/admin/flight-schedules/new')} variant="default"
			><Plus /> New flight schedule</Button
		>
	{/snippet}
	<FlightList flightSchedules={data.flightSchedules} />
</Page>
