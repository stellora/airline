<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import Group from 'lucide-svelte/icons/group'
	import Plane from 'lucide-svelte/icons/plane'
	import PlaneTakeoff from 'lucide-svelte/icons/plane-takeoff'
	import Settings from 'lucide-svelte/icons/settings'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'

	const { data, children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]', {
				params: { airlineSpec: data.airline.iataCode },
			})}><AirlineCode airline={data.airline} icon /></Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/manage/[airlineSpec]', {
				params: { airlineSpec: data.airline.iataCode },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Schedules',
			url: route('/manage/[airlineSpec]/schedules', {
				params: { airlineSpec: data.airline.iataCode },
			}),
			icon: CalendarRange,
		},
		{
			title: 'Flights',
			url: route('/manage/[airlineSpec]/flights', {
				params: { airlineSpec: data.airline.iataCode },
			}),
			icon: PlaneTakeoff,
		},
		{
			title: 'Fleets',
			url: route('/manage/[airlineSpec]/fleets', {
				params: { airlineSpec: data.airline.iataCode },
			}),
			icon: Group,
		},
		{
			title: 'Aircraft',
			url: route('/manage/[airlineSpec]/aircraft', {
				params: { airlineSpec: data.airline.iataCode },
			}),
			icon: Plane,
		},
	]}
>
	{#snippet breadcrumbActions()}
		<PageNavbarBreadcrumbActionsDropdownMenu>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<a
							{...props}
							href={route('/admin/airlines/[airlineSpec]', {
								params: { airlineSpec: page.params.airlineSpec },
							})}
						>
							<Settings /> Public info
						</a>
					{/snippet}
				</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/manage/[airlineSpec]', {
								params: { airlineSpec: page.params.airlineSpec },
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
</PageNav>

{@render children()}
