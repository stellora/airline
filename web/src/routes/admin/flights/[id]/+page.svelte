<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import FormattedDatetime from '$lib/components/formatted-datetime.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
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

<Page title={flightTitle(data.flight)}>
	{@const departureDateTime = parseZonedDateTime(data.flight.departureDateTime)}
	{@const arrivalDateTime = parseZonedDateTime(data.flight.arrivalDateTime)}

	<div class="flex flex-wrap-reverse gap-4 items-end">
		<Card.Root class="flex-grow-[1]">
			<Card.Header>
				<div class="flex gap-2 items-center">
					<ScheduleTitle
						schedule={data.flight}
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
						<AirportCode airport={data.flight.originAirport} link />
						<span class="text-muted-foreground">&mdash;</span>
						<FormattedDatetime value={departureDateTime}>
							{formatFlightDateTime(departureDateTime)}
						</FormattedDatetime>
					</DefinitionList.Item>
					<DefinitionList.Item title="Arrival">
						<AirportCode airport={data.flight.destinationAirport} link />
						<span class="text-muted-foreground">&mdash;</span>
						<FormattedDatetime value={arrivalDateTime}>
							{formatFlightTime(arrivalDateTime, {
								plusMinusDaysFrom: departureDateTime,
							})}
						</FormattedDatetime>
					</DefinitionList.Item>
					<DefinitionList.Item title="Duration">
						{formatFlightDuration(departureDateTime, arrivalDateTime)}
						<span class="text-muted-foreground">&mdash;</span>
						<Distance distanceMiles={data.flight.distanceMiles} />
					</DefinitionList.Item>
				</DefinitionList.Root>
			</Card.Content>
		</Card.Root>

		<Card.Root class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flight.originAirport, data.flight.destinationAirport]]}
				detailLevel="high"
			/>
		</Card.Root>
	</div>

	<Card.Root>
		<Card.Header>
			<Card.Title>Notes</Card.Title>
		</Card.Header>
		<Card.Content>
			{#if data.flight.notes}
				<p>{data.flight.notes}</p>
			{:else}
				<p class="text-sm text-muted-foreground">No notes yet.</p>
			{/if}
		</Card.Content>
	</Card.Root>
</Page>
