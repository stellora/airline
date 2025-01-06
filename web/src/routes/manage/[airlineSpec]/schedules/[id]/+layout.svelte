<script lang="ts">
	import { enhance } from '$app/forms'
	import { schema } from '$lib/airline.typebox'
	import { buttonVariants } from '$lib/components/ui/button/button.svelte'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Eye from 'lucide-svelte/icons/eye'
	import EyeOff from 'lucide-svelte/icons/eye-off'
	import PlaneTakeoff from 'lucide-svelte/icons/plane-takeoff'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import FlightScheduleForm from '../schedule-form.svelte'
	import FlightScheduleBreadcrumbItem from './flight-schedule-breadcrumb-item.svelte'

	const { children, data } = $props()
</script>

<BreadcrumbsForLayout>
	<FlightScheduleBreadcrumbItem flightSchedule={data.flightSchedule} />
</BreadcrumbsForLayout>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/manage/[airlineSpec]/schedules/[id]', {
				params: {
					airlineSpec: data.flightSchedule.airline.iataCode,
					id: data.flightSchedule.id.toString(),
				},
			}),
			icon: SquareMenu,
		},
		{
			title: 'Instances',
			url: route('/manage/[airlineSpec]/schedules/[id]/instances', {
				params: {
					airlineSpec: data.flightSchedule.airline.iataCode,
					id: data.flightSchedule.id.toString(),
				},
			}),
			icon: PlaneTakeoff,
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
							action={route('/manage/[airlineSpec]/schedules/[id]', {
								params: {
									airlineSpec: data.flightSchedule.airline.iataCode,
									id: data.flightSchedule.id.toString(),
								},
								query: '/setFlightSchedulePublished',
							})}
							use:enhance
							class="w-full [&>button]:w-full"
						>
							<input
								type="hidden"
								name="published"
								value={data.flightSchedule.published ? 'false' : 'true'}
							/>
							<button type="submit" {...props}>
								{#if data.flightSchedule.published}
									<EyeOff /> Unpublish
								{:else}
									<Eye /> Publish
								{/if}
							</button>
						</form>
					{/snippet}
				</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/manage/[airlineSpec]/schedules/[id]', {
								params: {
									airlineSpec: data.flightSchedule.airline.iataCode,
									id: data.flightSchedule.id.toString(),
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
	{/snippet}
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="edit-flight-schedule" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit flight schedule</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FlightScheduleForm
						action={route('/manage/[airlineSpec]/schedules/[id]', {
							params: {
								airlineSpec: data.flightSchedule.airline.iataCode,
								id: data.flightSchedule.id.toString(),
							},
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/flight-schedules/{id}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
