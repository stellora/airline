<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import Plus from 'lucide-svelte/icons/plus'
	import AirportForm from './airport-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="new-airport">
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Plus /> New airport</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>New airport</Dialog.Title>
				</Dialog.Header>
				<AirportForm action="?/create" submitLabel="Create" form={data.form} />
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Airports">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">IATA code</Table.Head>
					<Table.Head>Name</Table.Head>
					<Table.Head>Region</Table.Head>
					<Table.Head class="text-right"></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.airports && data.airports.length > 0}
				<Table.Body>
					{#each data.airports as airport (airport.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell class="font-bold text-lg">
								<AirportCode tooltip={false} {airport} />
							</Table.Cell>
							<Table.Cell>{airport.name}</Table.Cell>
							<Table.Cell>{airport.region}, {airport.country}</Table.Cell>
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/admin/airports/[airportSpec]', {
										params: { airportSpec: airport.iataCode },
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
				<Table.Caption class="mb-4">No airports found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
