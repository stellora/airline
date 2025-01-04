<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightScheduleForm from './flight-schedule-form.svelte'
	import FlightScheduleList from './flight-schedule-list.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="new-flight-schedule">
			<Dialog.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New flight schedule
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>New flight schedule</Dialog.Title>
				</Dialog.Header>
				<FlightScheduleForm action="?/create" submitLabel="Create" form={data.form} />
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Flight schedules">
	<FlightScheduleList flightSchedules={data.flightSchedules} />
</Page>
