<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightScheduleForm from './flight-instance-form.svelte'
	import FlightScheduleList from './flight-instance-list.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="new-flight-instance">
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Plus /> New flight instance</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>New flight instance</Dialog.Title>
				</Dialog.Header>
				<FlightScheduleForm action="?/create" submitLabel="Create" form={data.form} />
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Flight instances">
	<FlightScheduleList flightInstances={data.flightInstances} />
</Page>
