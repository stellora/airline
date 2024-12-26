<script lang="ts">
	import { page } from '$app/state'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import Boxes from 'lucide-svelte/icons/boxes'
	import House from 'lucide-svelte/icons/house'
	import Tag from 'lucide-svelte/icons/tag'

	type Item = {
		title: string
		url: string
		icon: typeof House
	}

	const items: Item[] = [
		{
			title: 'Home',
			url: '/',
			icon: House
		}
	]

	const adminItems: Item[] = [
		{
			title: 'Flights',
			url: '/admin/flights',
			icon: Boxes
		},
		{
			title: 'Airports',
			url: '/admin/airports',
			icon: Tag
		}
	]
</script>

<Sidebar.Root collapsible="offcanvas" variant="inset">
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>Airline</Sidebar.GroupLabel>
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
		<Sidebar.Group>
			<Sidebar.GroupLabel>Admin</Sidebar.GroupLabel>
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
	</Sidebar.Content>
</Sidebar.Root>
