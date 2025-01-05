<script lang="ts">
	import ItineraryTitle from '$lib/components/itinerary-title.svelte'
	import SeatMap from '$lib/components/seat-map/seat-map.svelte'
	import * as Card from '$lib/components/ui/card/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table/index.js'
	import { formatFlightDate } from '$lib/datetime-helpers'
	import { flightTitle } from '$lib/flight-helpers'
	import { parseZonedDateTime } from '@internationalized/date'

	let { data } = $props()
</script>

<Page
	title={`Passengers on ${flightTitle(data.flightInstance)} ${formatFlightDate(parseZonedDateTime(data.flightInstance.departureDateTime))}`}
>
	<Card.Root>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[150px]">Seat</Table.Head>
					<Table.Head class="w-[300px]">Passenger</Table.Head>
					<Table.Head class="w-[150px]">Itinerary</Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.seatAssignments && data.seatAssignments.length > 0}
				<Table.Body>
					{#each data.seatAssignments as a (a.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell>
								{a.seat}
							</Table.Cell>
							<Table.Cell>
								{a.passenger.name}
							</Table.Cell>
							<Table.Cell>
								<ItineraryTitle itinerary={a.itinerary} link class="font-bold" />
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			{:else}
				<Table.Caption class="mb-4">No seat assignments found</Table.Caption>
			{/if}
		</Table.Root>
	</Card.Root>

	<Card.Root class="self-start">
		<Card.Header>
			<Card.Title>Seat map</Card.Title>
		</Card.Header>
		<Card.CardContent>
			<SeatMap seatAssignments={data.seatAssignments} />
		</Card.CardContent>
	</Card.Root>
</Page>
