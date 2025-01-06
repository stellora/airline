<script lang="ts">
	import AircraftRegistration from '$lib/components/aircraft-registration.svelte'
	import FleetTitle from '$lib/components/fleet-title.svelte'
	import FlightStatus from '$lib/components/flight-status.svelte'
	import FormattedDatetime from '$lib/components/formatted-datetime.svelte'
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Table from '$lib/components/ui/table'
	import { formatFlightDate, formatFlightDuration, formatFlightTime } from '$lib/datetime-helpers'
	import { route } from '$lib/route-helpers'
	import type { Flight } from '$lib/types'
	import { parseZonedDateTime } from '@internationalized/date'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'

	let { flights, showFlightInfo }: { flights: Flight[]; showFlightInfo?: boolean } = $props()
</script>

<Card>
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-[155px]">Date</Table.Head>
				{#if showFlightInfo}
					<Table.Head class="w-[180px]">Flight</Table.Head>
				{/if}
				<Table.Head class="w-[95px]">Aircraft</Table.Head>
				<Table.Head class="w-[105px]">Flight time</Table.Head>
				<Table.Head class="w-[105px]">Departure</Table.Head>
				<Table.Head class="w-[105px]">Arrival</Table.Head>
				<Table.Head class="">Status</Table.Head>
				<Table.Head class="text-right" />
			</Table.Row>
		</Table.Header>
		{#if flights && flights.length > 0}
			<Table.Body>
				{#each flights as flight (flight.id)}
					{@const departureDateTime = parseZonedDateTime(flight.departureDateTime)}
					{@const arrivalDateTime = parseZonedDateTime(flight.arrivalDateTime)}
					<Table.Row class="stretched-link-container group">
						<Table.Cell>
							<FormattedDatetime value={departureDateTime} class="cursor-help">
								{formatFlightDate(departureDateTime)}
							</FormattedDatetime>
						</Table.Cell>
						{#if showFlightInfo}
							<Table.Cell><ScheduleTitle schedule={flight} /></Table.Cell>
						{/if}
						<Table.Cell
							><div class="inline-flex flex-col gap-1">
								{#if flight.aircraft}
									<AircraftRegistration aircraft={flight.aircraft} showAircraftType link />
								{:else}
									<FleetTitle fleet={flight.fleet} link />
								{/if}
							</div></Table.Cell
						>
						<Table.Cell>{formatFlightDuration(departureDateTime, arrivalDateTime)}</Table.Cell>
						<Table.Cell
							><FormattedDatetime value={departureDateTime}>
								{formatFlightTime(departureDateTime)}
							</FormattedDatetime></Table.Cell
						>
						<Table.Cell>
							<FormattedDatetime value={arrivalDateTime}>
								{formatFlightTime(arrivalDateTime, {
									plusMinusDaysFrom: departureDateTime,
								})}
							</FormattedDatetime>
						</Table.Cell>
						<Table.Cell><FlightStatus {flight} /></Table.Cell>
						<Table.Cell class="text-right">
							<Button
								variant="link"
								href={route('/manage/[airlineSpec]/flights/[id]', {
									params: {
										airlineSpec: flight.airline.iataCode,
										id: flight.id.toString(),
									},
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
