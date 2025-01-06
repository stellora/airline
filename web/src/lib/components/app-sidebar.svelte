<script lang="ts">
	import { page } from '$app/state'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import { route } from '$lib/route-helpers'
	import Book from 'lucide-svelte/icons/book'
	import Building from 'lucide-svelte/icons/building'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import CircleUser from 'lucide-svelte/icons/circle-user'
	import Group from 'lucide-svelte/icons/group'
	import Info from 'lucide-svelte/icons/info'
	import List from 'lucide-svelte/icons/list'
	import MapPin from 'lucide-svelte/icons/map-pin'
	import Plane from 'lucide-svelte/icons/plane'
	import PlaneTakeoff from 'lucide-svelte/icons/plane-takeoff'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import TicketsPlane from 'lucide-svelte/icons/tickets-plane'
	import User from 'lucide-svelte/icons/user'
	import Users from 'lucide-svelte/icons/users'
	import Waypoints from 'lucide-svelte/icons/waypoints'
	import type { LayoutData } from '../../routes/$types'
	import AppSidebarAirlineSwitcher from './app-sidebar-airline-switcher.svelte'

	type Item = {
		title: string
		url: string
		icon: typeof TicketsPlane
	}

	const items: Item[] = [
		{
			title: 'Book flights',
			url: route('/'),
			icon: TicketsPlane,
		},
		{
			title: 'My bookings',
			url: '/bookings',
			icon: User,
		},
		{
			title: 'Flight status',
			url: '/flight-status',
			icon: Info,
		},
		{
			title: 'Routes',
			url: route('/admin/routes'),
			icon: Waypoints,
		},
	]

	const airlineAdminItems: (airlineIataCode: string) => Item[] = (airlineIataCode) => [
		{
			title: 'Overview',
			url: route('/manage/[airlineSpec]', {
				params: { airlineSpec: airlineIataCode },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Schedules',
			url: route('/manage/[airlineSpec]/schedules', {
				params: { airlineSpec: airlineIataCode },
			}),
			icon: CalendarRange,
		},
		{
			title: 'Fleets',
			url: route('/manage/[airlineSpec]/fleets', {
				params: { airlineSpec: airlineIataCode },
			}),
			icon: Group,
		},
		{
			title: 'Flights',
			url: route('/manage/[airlineSpec]/flights', {
				params: { airlineSpec: airlineIataCode },
			}),
			icon: PlaneTakeoff,
		},
		{
			title: 'Aircraft',
			url: route('/manage/[airlineSpec]/aircraft', {
				params: { airlineSpec: airlineIataCode },
			}),
			icon: Plane,
		},
	]

	const globalAdminItems: Item[] = [
		{
			title: 'Airports',
			url: route('/admin/airports'),
			icon: MapPin,
		},
		{
			title: 'Itineraries',
			url: '/admin/itineraries',
			icon: TicketsPlane,
		},
		{
			title: 'Passengers',
			url: '/admin/passengers',
			icon: Users,
		},
		{
			title: 'Airlines',
			url: route('/admin/airlines'),
			icon: Building,
		},
		{
			title: 'Aircraft types',
			url: route('/admin/aircraft-types'),
			icon: List,
		},
	]

	const devItems: Item[] = [
		{
			title: 'Storybook',
			url: route('/dev/storybook'),
			icon: Book,
		},
	]

	const layoutData = page.data as unknown as LayoutData

	const activeAirlineIataCode = $derived(
		page.url.pathname.startsWith('/manage/') ? page.params.airlineSpec : undefined,
	)
</script>

<Sidebar.Root collapsible="offcanvas" variant="sidebar">
	<Sidebar.Content>
		<Sidebar.Group>
			<AppSidebarAirlineSwitcher allAirlines={layoutData.allAirlines} {activeAirlineIataCode} />
			{#if activeAirlineIataCode}
				<Sidebar.GroupContent class="mt-2">
					<Sidebar.Menu>
						{#each airlineAdminItems(activeAirlineIataCode) as item (item.title)}
							<Sidebar.MenuItem>
								<Sidebar.MenuButton isActive={page.url.pathname === item.url}>
									{#snippet child({ props })}
										<a href={item.url} {...props}>
											<item.icon />
											<span>{item.title}</span>
										</a>
									{/snippet}
								</Sidebar.MenuButton>
							</Sidebar.MenuItem>
						{/each}
					</Sidebar.Menu>
				</Sidebar.GroupContent>
			{/if}
		</Sidebar.Group>
		<Sidebar.Separator />
		<Sidebar.Group>
			<Sidebar.GroupLabel>Global admin</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each globalAdminItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={page.url.pathname === item.url}>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
		<Sidebar.Separator />
		<Sidebar.Group>
			<Sidebar.GroupLabel>Public</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each items as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={page.url.pathname === item.url}>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
		<Sidebar.Separator />
		<Sidebar.Group>
			<Sidebar.GroupLabel>Dev</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each devItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={page.url.pathname === item.url}>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
	<Sidebar.Separator />
	<Sidebar.Footer>
		<div class="text-muted-foreground text-sm flex items-center gap-1.5">
			<CircleUser size="2.5ch" /> <strong>sqs</strong>
		</div></Sidebar.Footer
	>
</Sidebar.Root>
