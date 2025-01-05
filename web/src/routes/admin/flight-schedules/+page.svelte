<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightScheduleForm from './flight-schedule-form.svelte'
	import FlightScheduleList from './flight-schedule-list.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-flight-schedule" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New flight schedule
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New flight schedule</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FlightScheduleForm action="?/create" submitLabel="Create" form={data.form} />
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link href={route('/admin/flight-schedules')}>Schedules</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<Page title="Flight schedules">
	<FlightScheduleList flightSchedules={data.flightSchedules} />
</Page>
