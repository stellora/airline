<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Table from '$lib/components/ui/table'
	import AirportForm from './airport-form.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Airports</h1>
	<AirportForm {form} />
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
									href={`/admin/airports/${airport.id}`}
									class="stretched-link h-auto p-1 opacity-35 group-hover:opacity-100"
								>
									Manage
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
</div>
