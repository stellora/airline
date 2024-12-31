<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import FlightSparkRoute from '$lib/components/flight-spark-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Table from '$lib/components/ui/table'
	import type { Flight } from '$lib/types'

	let { flights }: { flights: Flight[] } = $props()
</script>

<Card>
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-[200px]">Flight number</Table.Head>
				<Table.Head>Route</Table.Head>
				<Table.Head class="text-right"></Table.Head>
			</Table.Row>
		</Table.Header>
		{#if flights && flights.length > 0}
			<Table.Body>
				{#each flights as flight (flight.id)}
					<Table.Row class="stretched-link-container group">
						<Table.Cell class="font-bold text-lg">
							{flight.number}
						</Table.Cell>
						<Table.Cell class="flex gap-2"
							><span
								><AirportCode airport={flight.originAirport} />&ndash;<AirportCode
									airport={flight.destinationAirport}
								/></span
							>
							<span class="text-muted-foreground">
								<Distance distanceMiles={flight.distanceMiles} /></span
							>
							<div class="border rounded">
								<FlightSparkRoute {flight} width={100} height={25} />
							</div>
						</Table.Cell>
						<Table.Cell class="text-right">
							<Button
								variant="link"
								href={`/admin/flights/${flight.id}`}
								class="stretched-link h-auto p-1 opacity-35 group-hover:opacity-100"
							>
								Manage
							</Button>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		{:else}
			<Table.Caption class="mb-4">No flights found</Table.Caption>
		{/if}
	</Table.Root>
</Card>
