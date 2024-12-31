<script lang="ts">
	import { page } from '$app/state'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import Book from 'lucide-svelte/icons/book'
	import CircleUser from 'lucide-svelte/icons/circle-user'
	import Info from 'lucide-svelte/icons/info'
	import MapPin from 'lucide-svelte/icons/map-pin'
	import Plane from 'lucide-svelte/icons/plane'
	import Spline from 'lucide-svelte/icons/spline'
	import TicketsPlane from 'lucide-svelte/icons/tickets-plane'
	import User from 'lucide-svelte/icons/user'
	import Users from 'lucide-svelte/icons/users'
	import Waypoints from 'lucide-svelte/icons/waypoints'

	type Item = {
		title: string
		url: string
		icon: typeof TicketsPlane
	}

	const items: Item[] = [
		{
			title: 'Book flights',
			url: '/',
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
	]

	const adminItems: Item[] = [
		{
			title: 'Flights',
			url: '/admin/flights',
			icon: Spline,
		},
		{
			title: 'Routes',
			url: '/admin/routes',
			icon: Waypoints,
		},
		{
			title: 'Airports',
			url: '/admin/airports',
			icon: MapPin,
		},
		{
			title: 'Aircraft',
			url: '/admin/aircraft',
			icon: Plane,
		},
		{
			title: 'Passengers',
			url: '/admin/passengers',
			icon: Users,
		},
	]

	const devItems: Item[] = [
		{
			title: 'Storybook',
			url: '/dev/storybook',
			icon: Book,
		},
	]
</script>

<Sidebar.Root collapsible="offcanvas" variant="sidebar">
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>Stellora Airlines</Sidebar.GroupLabel>
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
			<Sidebar.GroupLabel>Administration</Sidebar.GroupLabel>
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
