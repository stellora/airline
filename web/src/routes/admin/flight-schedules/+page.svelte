<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Sheet from '$lib/components/ui/sheet/index.js'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightScheduleForm from './flight-schedule-form.svelte'
	import FlightScheduleList from './flight-schedule-list.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Sheet.SheetByNavigationState id="new-flight-schedule">
			<Sheet.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New flight schedule
			</Sheet.Trigger>
			<Sheet.Content>
				<Sheet.Header>
					<Sheet.Title>New flight schedule</Sheet.Title>
				</Sheet.Header>
				<FlightScheduleForm action="?/create" submitLabel="Create" form={data.form} />
			</Sheet.Content>
		</Sheet.SheetByNavigationState>
	{/snippet}
</PageNav>

<Page title="Flight schedules">
	<FlightScheduleList flightSchedules={data.flightSchedules} />
</Page>
