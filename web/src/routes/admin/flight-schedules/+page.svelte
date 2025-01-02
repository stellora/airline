<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import { Button } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import { PAGE_NAVBAR_ACTIONS_ID } from '$lib/components/ui/page/page-navbar.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'
	import NewFlightScheduleDialogContent from './new-flight-schedule-dialog-content.svelte'

	let { data } = $props()
</script>

<Portal target={PAGE_NAVBAR_ACTIONS_ID}>
	<div class="flex justify-end">
		<Dialog.RootByNavigationState id="foo">
			<Dialog.Trigger>
				<Button variant="secondary" size="sm" class="h-[unset] py-1 px-2"
					><Plus /> New flight schedule</Button
				>
			</Dialog.Trigger>
			<NewFlightScheduleDialogContent action="?/create" form={data.form} />
		</Dialog.RootByNavigationState>
	</div>
</Portal>

<Page title="Flight schedules">
	<FlightList flightSchedules={data.flightSchedules} />
</Page>
