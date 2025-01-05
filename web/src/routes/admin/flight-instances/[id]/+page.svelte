<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Card, CardHeader, CardTitle } from '$lib/components/ui/card'
	import CardContent from '$lib/components/ui/card/card-content.svelte'
	import * as DefinitionList from '$lib/components/ui/definition-list/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import { formatFlightDate, formatFlightTime } from '$lib/datetime-helpers.js'
	import { flightTitle } from '$lib/flight-helpers'
	import { parseZonedDateTime } from '@internationalized/date'

	let { data } = $props()
</script>

<Page
	title={`${flightTitle(data.flightInstance)} flight on ${data.flightInstance.scheduleInstanceDate}`}
	showTitleHeading={true}
>
	{#snippet titleElement(className)}
		<div class="flex items-baseline gap-2">
			<FlightTitle
				flight={data.flightInstance}
				class={className}
				subtitleClass="text-base"
				as="h1"
			/>
			<span class="text-4xl">{data.flightInstance.scheduleInstanceDate}</span>
		</div>
	{/snippet}

	{@const departureDateTime = parseZonedDateTime(data.flightInstance.departureDateTime)}
	{@const arrivalDateTime = parseZonedDateTime(data.flightInstance.arrivalDateTime)}

	<div class="flex flex-wrap-reverse gap-4">
		<Card class="flex-grow-[1]">
			<CardHeader>
				<CardTitle>Flight details</CardTitle>
			</CardHeader>
			<CardContent>
				<DefinitionList.Root>
					<DefinitionList.Item title="Departure">
						<AirportCode airport={data.flightInstance.originAirport} />
						<span class="text-muted-foreground">&mdash;</span>
						{formatFlightDate(departureDateTime)}
						{formatFlightTime(departureDateTime)}
					</DefinitionList.Item>
					<DefinitionList.Item title="Arrival">
						<AirportCode airport={data.flightInstance.destinationAirport} />
						<span class="text-muted-foreground">&mdash;</span>
						{formatFlightDate(arrivalDateTime)}
						{formatFlightTime(arrivalDateTime, {
							plusMinusDaysFrom: departureDateTime,
						})}
					</DefinitionList.Item>
				</DefinitionList.Root>
			</CardContent>
		</Card>

		<Card class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flightInstance.originAirport, data.flightInstance.destinationAirport]]}
				detailLevel="high"
			/>
		</Card>
	</div>
</Page>
