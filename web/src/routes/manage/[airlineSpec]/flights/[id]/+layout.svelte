<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { formatFlightDate } from '$lib/datetime-helpers'
	import { route } from '$lib/route-helpers'
	import { cn } from '$lib/utils'
	import { parseZonedDateTime } from '@internationalized/date'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import Grid3x3 from 'lucide-svelte/icons/grid-3x3'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import Users from 'lucide-svelte/icons/users'
	import type { ClassNameValue } from 'tailwind-merge'
	import ScheduleBreadcrumbItem from '../../schedules/[id]/schedule-breadcrumb-item.svelte'
	import FlightForm from './flight-form.svelte'

	const { children, data } = $props()
</script>

<BreadcrumbsForLayout restart>
	{#if data.schedule}
		<ScheduleBreadcrumbItem schedule={data.schedule} />
		<Breadcrumb.Separator />
	{/if}
	<Breadcrumb.Item
		><Breadcrumb.Link
			href={route('/manage/[airlineSpec]/flights/[id]', {
				params: { airlineSpec: data.flight.airline.iataCode, id: data.flight.id.toString() },
			})}
			>Flight on {formatFlightDate(
				parseZonedDateTime(data.flight.departureDateTime),
			)}</Breadcrumb.Link
		></Breadcrumb.Item
	>
</BreadcrumbsForLayout>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/manage/[airlineSpec]/flights/[id]', {
				params: { airlineSpec: page.params.airlineSpec, id: page.params.id },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Passengers',
			url: route('/manage/[airlineSpec]/flights/[id]/passengers', {
				params: { airlineSpec: page.params.airlineSpec, id: page.params.id },
			}),
			icon: Users,
		},
		{
			title: 'Seat map',
			url: route('/manage/[airlineSpec]/flights/[id]/seat-map', {
				params: { airlineSpec: page.params.airlineSpec, id: page.params.id },
			}),
			icon: Grid3x3,
		},
	]}
>
	{#snippet breadcrumbActions()}
		<Drawer.DrawerByNavigationState id="edit-flight" direction="right">
			<PageNavbarBreadcrumbActionsDropdownMenu>
				<DropdownMenu.Group>
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<Drawer.Trigger {...props} class={cn(props.class as ClassNameValue, 'w-full')}>
								<Settings2 /> Edit
							</Drawer.Trigger>
						{/snippet}
					</DropdownMenu.Item>
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<form
								method="POST"
								action={route('/manage/[airlineSpec]/flights/[id]', {
									params: {
										airlineSpec: page.params.airlineSpec,
										id: page.params.id,
									},
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
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit flight</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FlightForm
						flight={data.flight}
						action={route('/manage/[airlineSpec]/flights/[id]', {
							params: { airlineSpec: data.flight.airline.iataCode, id: data.flight.id.toString() },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/flights/{id}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
	{#snippet actions()}
		{#if data.flight.scheduleID}
			<Button
				variant="outline"
				size="pageNavbar"
				href={route('/manage/[airlineSpec]/schedules/[id]', {
					params: {
						airlineSpec: data.flight.airline.iataCode,
						id: data.flight.scheduleID.toString(),
					},
				})}><CalendarRange /> View schedule</Button
			>
		{/if}
	{/snippet}
</PageNav>

{@render children()}
