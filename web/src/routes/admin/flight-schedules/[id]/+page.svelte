<script lang="ts">
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import * as Card from '$lib/components/ui/card'
	import * as DefinitionList from '$lib/components/ui/definition-list/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import { formatFlightDuration } from '$lib/datetime-helpers.js'
	import { flightTitle, formatDaysOfWeek } from '$lib/flight-helpers'

	let { data } = $props()
</script>

<Page title={`${flightTitle(data.flightSchedule)} schedule`}>
	<div class="flex flex-wrap-reverse gap-4 items-end">
		<Card.Root class="flex-grow-[1]">
			<Card.Header>
				<FlightTitle
					flight={data.flightSchedule}
					class="text-4xl font-bold"
					subtitleClass="text-base"
					as="h1"
					showRoute={false}
				/>
			</Card.Header>
			<Card.Content>
				<DefinitionList.Root>
					<DefinitionList.Item title="Route">
						<AirportCode
							airport={data.flightSchedule.originAirport}
							tooltip
							link
						/>&ndash;<AirportCode airport={data.flightSchedule.destinationAirport} tooltip link />
						<span class="text-muted-foreground text-sm"
							><Distance distanceMiles={data.flightSchedule.distanceMiles} /></span
						>
					</DefinitionList.Item>
					<DefinitionList.Item title="Operating airline">
						<AirlineCode
							airline={data.flightSchedule.airline}
							icon={false}
							tooltip={false}
							link
							showName
						/>
					</DefinitionList.Item>
					<DefinitionList.Item title="Aircraft type">
						<AircraftTypeCode aircraftType={data.flightSchedule.aircraftType} showName />
					</DefinitionList.Item>
					<DefinitionList.Item title="Scheduled">
						{data.flightSchedule.startDate} to {data.flightSchedule.endDate}
					</DefinitionList.Item>
					<DefinitionList.Item title="Timings">
						{data.flightSchedule.departureTime}
						{formatDaysOfWeek(data.flightSchedule.daysOfWeek)} ({formatFlightDuration(
							data.flightSchedule.durationSec,
						)})
					</DefinitionList.Item>
				</DefinitionList.Root>
			</Card.Content>
		</Card.Root>

		<Card.Root class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flightSchedule.originAirport, data.flightSchedule.destinationAirport]]}
				detailLevel="high"
			/>
		</Card.Root>
	</div>
</Page>
