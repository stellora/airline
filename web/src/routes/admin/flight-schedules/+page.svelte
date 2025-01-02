<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightList from './flight-list.svelte'
	import FlightScheduleForm from './flight-schedule-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="foo">
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Plus /> New flight schedule</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>New flight schedule</Dialog.Title>
				</Dialog.Header>
				<FlightScheduleForm action="?/create" form={data.form} />
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Flight schedules">
	<FlightList flightSchedules={data.flightSchedules} />
</Page>
