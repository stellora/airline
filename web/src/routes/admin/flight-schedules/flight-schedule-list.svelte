<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import FlightSparkRoute from '$lib/components/flight-spark-route.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import type { FlightSchedule } from '$lib/types'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'

	let { flightSchedules }: { flightSchedules: FlightSchedule[] } = $props()
</script>

<Card>
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-[200px]">Flight number</Table.Head>
				<Table.Head>Route</Table.Head>
				<Table.Head class="text-right" />
			</Table.Row>
		</Table.Header>
		{#if flightSchedules && flightSchedules.length > 0}
			<Table.Body>
				{#each flightSchedules as flight (flight.id)}
					<Table.Row class="stretched-link-container group">
						<Table.Cell class="font-bold text-lg">
							<FlightTitle {flight} as="span" showRoute={false} />
						</Table.Cell>
						<Table.Cell class="flex items-center gap-2"
							><div class="inline-flex flex-col gap-1">
								<span
									><AirportCode airport={flight.originAirport} />&ndash;<AirportCode
										airport={flight.destinationAirport}
									/></span
								>
								<span class="text-muted-foreground text-xs">
									<Distance distanceMiles={flight.distanceMiles} /></span
								>
							</div>
							<div class="border rounded">
								<FlightSparkRoute {flight} width={100} height={36} />
							</div>
						</Table.Cell>
						<Table.Cell class="text-right">
							<Button
								variant="link"
								href={route('/admin/flight-schedules/[id]', {
									params: { id: flight.id.toString() },
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
			<Table.Caption class="mb-4">No flights found</Table.Caption>
		{/if}
	</Table.Root>
</Card>
