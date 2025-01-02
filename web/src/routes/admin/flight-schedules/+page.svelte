<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'
	import NewFlightScheduleDialogContent from './new-flight-schedule-dialog-content.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="foo">
			<Dialog.Trigger>
				<Button variant="secondary" size="sm" class="h-[unset] py-1 px-2"
					><Plus /> New flight schedule</Button
				>
			</Dialog.Trigger>
			<NewFlightScheduleDialogContent action="?/create" form={data.form} />
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Flight schedules">
	<FlightList flightSchedules={data.flightSchedules} />
</Page>
