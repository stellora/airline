<script lang="ts">
	import { page } from '$app/state'
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Table from '$lib/components/ui/table'
	import { formatterForFlightDateTime } from '$lib/datetime-helpers'
	import { route } from '$lib/route-helpers'
	import type { FlightInstance } from '$lib/types'
	import { parseDateTime } from '@internationalized/date'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'

	let { flightInstances }: { flightInstances: FlightInstance[] } = $props()
</script>

<FlightTitle flight={page.data.flightSchedule} as="span" showRoute={true} />

<Card>
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-[125px]">Departure</Table.Head>
				<Table.Head class="w-[125px]">Arrival</Table.Head>
				<Table.Head>Aircraft</Table.Head>
				<Table.Head class="text-right" />
			</Table.Row>
		</Table.Header>
		{#if flightInstances && flightInstances.length > 0}
			<Table.Body>
				{#each flightInstances as flight (flight.id)}
					{@const departureDateTime = parseDateTime(flight.departureDateTime).toDate('UTC')}
					{@const arrivalDateTime = parseDateTime(flight.arrivalDateTime).toDate('UTC')}
					<Table.Row class="stretched-link-container group">
						<Table.Cell>
							{formatterForFlightDateTime(departureDateTime).format(departureDateTime)}</Table.Cell
						>
						<Table.Cell>{flight.arrivalDateTime}</Table.Cell>
						<Table.Cell
							><div class="inline-flex flex-col gap-1">
								<AircraftTypeCode aircraftType={flight.aircraftType} />
							</div></Table.Cell
						>
						<Table.Cell class="text-right">
							<Button
								variant="link"
								href={route('/admin/flight-instances/[id]', {
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
			<Table.Caption class="mb-4">No flight instances found</Table.Caption>
		{/if}
	</Table.Root>
</Card>
