<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import ItineraryTitle from '$lib/components/itinerary-title.svelte'
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { flightTitle } from '$lib/flight-helpers'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import Plus from 'lucide-svelte/icons/plus'
	import ItineraryForm from './create-itinerary-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-itinerary" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New itinerary
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New itinerary</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<ItineraryForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/itineraries']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<Page title="Itineraries">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[150px]">Record locator</Table.Head>
					<Table.Head class="w-[30%]">Passengers</Table.Head>
					<Table.Head class="w-[30%]">Flights</Table.Head>
					<Table.Head class="text-right"></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.itineraries && data.itineraries.length > 0}
				<Table.Body>
					{#each data.itineraries as itinerary (itinerary.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell>
								<ItineraryTitle {itinerary} class="font-bold" />
							</Table.Cell>
							<Table.Cell>
								<ul>
									{#each itinerary.passengers as passenger (passenger.id)}
										<li>{passenger.name}</li>
									{/each}
								</ul>
							</Table.Cell>
							<Table.Cell>
								<ul>
									{#each itinerary.flights as flight (flight.id)}
										<li>{flightTitle(flight)}</li>
									{/each}
								</ul>
							</Table.Cell>
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/admin/itineraries/[itinerarySpec]', {
										params: { itinerarySpec: itinerary.recordID },
									})}
									class="stretched-link h-auto p-1 opacity-35 group-hover:opacity-100"
								>
									<ChevronRight />
								</Button>
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			{:else}
				<Table.Caption class="mb-4">No itineraries found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
