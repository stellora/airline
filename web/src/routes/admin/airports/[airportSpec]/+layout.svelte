<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import AirportCode from '$lib/components/airport-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button/button.svelte'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import AirportForm from '../airport-form.svelte'
	import { schema } from '$lib/airline.typebox'

	const { children, data } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/admin/airports/[airportSpec]', {
				params: { airportSpec: data.airport.iataCode },
			})}
			><AirportCode airport={data.airport} />
		</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/admin/airports/[airportSpec]', {
				params: { airportSpec: page.params.airportSpec },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Schedule',
			url: route('/admin/airports/[airportSpec]/flights', {
				params: { airportSpec: page.params.airportSpec },
			}),
			icon: CalendarRange,
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
							action={route('/admin/airports/[airportSpec]', {
								params: { airportSpec: page.params.airportSpec },
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
		<Drawer.DrawerByNavigationState id="edit-airport" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit airport</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<AirportForm
						action={route('/admin/airports/[airportSpec]', {
							params: { airportSpec: page.params.airportSpec },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/airports/{airportSpec}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
