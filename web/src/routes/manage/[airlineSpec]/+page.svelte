<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import * as Card from '$lib/components/ui/card'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import Settings from 'lucide-svelte/icons/settings'
	import Trash from 'lucide-svelte/icons/trash'

	let { data } = $props()
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

<PageNav>
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

<Page title={`${data.airline.iataCode}: ${data.airline.name}`}>
	<Card.Root>
		<Card.Header>
			<Card.Title><AirlineCode airline={data.airline} icon /></Card.Title>
			<Card.Description>
				{data.airline.name}
			</Card.Description>
		</Card.Header>
		<Card.Content />
	</Card.Root>

	{#await data.schedules then flights}
		{#if flights && flights.length > 0}
			<Card.Root class="overflow-hidden">
				<Card.Header>
					<Card.Title>Route map</Card.Title>
				</Card.Header>
				<GreatCircleRoute
					routes={flights.map((flight) => [flight.originAirport, flight.destinationAirport])}
				/>
			</Card.Root>
		{/if}
	{/await}
</Page>
