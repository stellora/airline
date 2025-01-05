<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import AirportCode from '$lib/components/airport-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
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
			title: 'Flight schedule',
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
		<Dialog.RootByNavigationState id="edit-airport">
			<Dialog.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>Edit airport</Dialog.Title>
				</Dialog.Header>
				<AirportForm
					action={route('/admin/airports/[airportSpec]', {
						params: { airportSpec: page.params.airportSpec },
						query: '/update',
					})}
					submitLabel="Save"
					form={data.form}
				/>
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
