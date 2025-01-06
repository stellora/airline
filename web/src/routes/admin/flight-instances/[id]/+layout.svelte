<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import Button, { buttonVariants } from '$lib/components/ui/button/button.svelte'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { formatFlightDate } from '$lib/datetime-helpers'
	import { route } from '$lib/route-helpers'
	import { parseZonedDateTime } from '@internationalized/date'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import Grid3x3 from 'lucide-svelte/icons/grid-3x3'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import Users from 'lucide-svelte/icons/users'
	import ScheduleBreadcrumbItem from '../../../manage/[airlineSpec]/schedules/[id]/flight-schedule-breadcrumb-item.svelte'
	import FlightInstanceForm from './flight-instance-form.svelte'

	const { children, data } = $props()
</script>

<BreadcrumbsForLayout restart>
	{#if data.schedule}
		<ScheduleBreadcrumbItem schedule={data.schedule} />
		<Breadcrumb.Separator />
	{/if}
	<Breadcrumb.Item
		><Breadcrumb.Link
			href={route('/admin/flight-instances/[id]', {
				params: { id: data.flightInstance.id.toString() },
			})}
			>Flight on {formatFlightDate(
				parseZonedDateTime(data.flightInstance.departureDateTime),
			)}</Breadcrumb.Link
		></Breadcrumb.Item
	>
</BreadcrumbsForLayout>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/admin/flight-instances/[id]', {
				params: { id: page.params.id },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Passengers',
			url: route('/admin/flight-instances/[id]/passengers', {
				params: { id: page.params.id },
			}),
			icon: Users,
		},
		{
			title: 'Seat map',
			url: route('/admin/flight-instances/[id]/seat-map', {
				params: { id: page.params.id },
			}),
			icon: Grid3x3,
		},
	]}
>
	{#snippet breadcrumbActions()}
		<PageNavbarBreadcrumbActionsDropdownMenu>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/flight-instances/[id]', {
								params: { id: page.params.id },
								query: '/delete',
							})}
							use:enhance={({ cancel }) => {
								if (!confirm('Really delete?')) {
									cancel()
								}
							}}
							class="w-full [&>button]:w-full"
						>
							<button type="submit" {...props}>
								<Trash /> Delete...
							</button>
						</form>
					{/snippet}</DropdownMenu.Item
				>
			</DropdownMenu.Group>
		</PageNavbarBreadcrumbActionsDropdownMenu>
	{/snippet}
	{#snippet actions()}
		{#if data.flightInstance.scheduleID}
			<Button
				variant="outline"
				size="pageNavbar"
				href={route('/manage/[airlineSpec]/schedules/[id]', {
					params: {
						airlineSpec: data.flightInstance.airline.iataCode,
						id: data.flightInstance.scheduleID.toString(),
					},
				})}><CalendarRange /> View schedule</Button
			>
		{/if}
		<Drawer.DrawerByNavigationState id="edit-flight-instance" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit flight instance</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FlightInstanceForm
						flightInstance={data.flightInstance}
						action={route('/admin/flight-instances/[id]', {
							params: { id: page.params.id },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/flight-instances/{id}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
