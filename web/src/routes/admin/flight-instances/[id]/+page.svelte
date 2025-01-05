<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import FormattedDatetime from '$lib/components/formatted-datetime.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Badge } from '$lib/components/ui/badge'
	import * as Card from '$lib/components/ui/card'
	import * as DefinitionList from '$lib/components/ui/definition-list/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import {
		formatFlightDate,
		formatFlightDateTime,
		formatFlightDuration,
		formatFlightTime,
	} from '$lib/datetime-helpers.js'
	import { flightTitle } from '$lib/flight-helpers'
	import { parseZonedDateTime } from '@internationalized/date'
	import CalendarDays from 'lucide-svelte/icons/calendar-days'

	let { data } = $props()
</script>

<Page
	title={`${flightTitle(data.flightInstance)} flight on ${data.flightInstance.scheduleInstanceDate}`}
>
	{@const departureDateTime = parseZonedDateTime(data.flightInstance.departureDateTime)}
	{@const arrivalDateTime = parseZonedDateTime(data.flightInstance.arrivalDateTime)}

	<div class="flex flex-wrap-reverse gap-4">
		<Card.Root class="flex-grow-[1]">
			<Card.Header>
				<div class="flex gap-2 items-center">
					<FlightTitle
						flight={data.flightInstance}
						class="text-4xl font-bold"
						subtitleClass="text-base"
						as="h1"
						showRoute={false}
					/>
					<Badge variant="secondary" size="xl" class="self-start"
						><CalendarDays />
						{formatFlightDate(departureDateTime)}</Badge
					>
				</div>
			</Card.Header>
			<Card.Content>
				<DefinitionList.Root>
					<DefinitionList.Item title="Departure">
						<AirportCode airport={data.flightInstance.originAirport} link />
						<span class="text-muted-foreground">&mdash;</span>
						<FormattedDatetime value={departureDateTime}>
							{formatFlightDateTime(departureDateTime)}
						</FormattedDatetime>
					</DefinitionList.Item>
					<DefinitionList.Item title="Arrival">
						<AirportCode airport={data.flightInstance.destinationAirport} link />
						<span class="text-muted-foreground">&mdash;</span>
						<FormattedDatetime value={arrivalDateTime}>
							{formatFlightTime(arrivalDateTime, {
								plusMinusDaysFrom: departureDateTime,
							})}
						</FormattedDatetime>
					</DefinitionList.Item>
					<DefinitionList.Item title="Duration">
						{formatFlightDuration(
							parseZonedDateTime(data.flightInstance.departureDateTime),
							parseZonedDateTime(data.flightInstance.arrivalDateTime),
						)},
						<Distance distanceMiles={data.flightInstance.distanceMiles} />
					</DefinitionList.Item>
				</DefinitionList.Root>
			</Card.Content>
		</Card.Root>

		<Card.Root class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flightInstance.originAirport, data.flightInstance.destinationAirport]]}
				detailLevel="high"
			/>
		</Card.Root>
	</div>
</Page>
