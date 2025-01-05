<script lang="ts">
	import { page } from '$app/state'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import { route } from '$lib/route-helpers'
	import Book from 'lucide-svelte/icons/book'
	import Building from 'lucide-svelte/icons/building'
	import CalendarRange from 'lucide-svelte/icons/calendar-range'
	import CircleUser from 'lucide-svelte/icons/circle-user'
	import Info from 'lucide-svelte/icons/info'
	import List from 'lucide-svelte/icons/list'
	import MapPin from 'lucide-svelte/icons/map-pin'
	import Plane from 'lucide-svelte/icons/plane'
	import PlaneTakeoff from 'lucide-svelte/icons/plane-takeoff'
	import TicketsPlane from 'lucide-svelte/icons/tickets-plane'
	import User from 'lucide-svelte/icons/user'
	import Users from 'lucide-svelte/icons/users'
	import Waypoints from 'lucide-svelte/icons/waypoints'
	import { buttonVariants } from './ui/button'

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

	const adminItems: Item[] = [
		{
			title: 'Schedules',
			url: route('/admin/flight-schedules'),
			icon: CalendarRange,
		},
		{
			title: 'Flights',
			url: route('/admin/flight-instances'),
			icon: PlaneTakeoff,
		},
		{
			title: 'Aircraft',
			url: '/admin/aircraft',
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
</script>

<Sidebar.Root collapsible="offcanvas" variant="sidebar">
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>
				Stellora Airlines
				<Sidebar.Trigger
					location="sidebar"
					class={buttonVariants({
						variant: 'ghost',
						size: 'pageNavbar',
						class: '[&>svg]:!size-4 px-2 -mr-2 ml-auto',
					})}
				/>
			</Sidebar.GroupLabel>
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
			<Sidebar.GroupLabel>Airline</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each adminItems as item (item.title)}
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
