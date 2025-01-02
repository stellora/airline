<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import { Button } from '$lib/components/ui/button'
	import { PAGE_NAVBAR_ACTIONS_ID } from '$lib/components/ui/page/page-navbar.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'

	let { data } = $props()
</script>

<Portal target={PAGE_NAVBAR_ACTIONS_ID}>
	<div class="flex justify-end">
		<Button
			href={route('/admin/flight-schedules/new')}
			variant="secondary"
			size="sm"
			class="h-[unset] py-1 px-2"><Plus /> New flight schedule</Button
		>
	</div>
</Portal>

<Page title="Flight schedules">
	<FlightList flightSchedules={data.flightSchedules} />
</Page>
