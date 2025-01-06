<script lang="ts">
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import ScheduleForm from './schedule-form.svelte'
	import ScheduleList from './schedule-list.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-schedule" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New schedule
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New schedule</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<ScheduleForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/schedules']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]/schedules', {
				params: { airlineSpec: page.params.airlineSpec },
			})}>Schedules</Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<Page title={`Schedules - ${data.airline.iataCode}`}>
	<ScheduleList schedules={data.schedules} />
</Page>
