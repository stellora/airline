<script lang="ts">
	import AirportCode from '$lib/components/airport-code.svelte'
	import FlightStatus from '$lib/components/flight-status.svelte'
	import FormattedDatetime from '$lib/components/formatted-datetime.svelte'
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Card from '$lib/components/ui/card/index.js'
	import * as DefinitionList from '$lib/components/ui/definition-list/index.js'
	import { formatFlightDate, formatFlightDuration, formatFlightTime } from '$lib/datetime-helpers'
	import type { Flight } from '$lib/types'
	import { parseZonedDateTime } from '@internationalized/date'
	import Armchair from 'lucide-svelte/icons/armchair'
	import CircleCheckBig from 'lucide-svelte/icons/circle-check-big'
	import Plane from 'lucide-svelte/icons/plane'

	const { flight }: { flight: Flight } = $props()

	let departureDateTime = $derived(parseZonedDateTime(flight.departureDateTime))
	let arrivalDateTime = $derived(parseZonedDateTime(flight.arrivalDateTime))
</script>

<Card.Root>
	<Card.Content class="flex gap-6 p-0 [&>*]:p-4">
		<div class="flex-1 flex flex-col gap-4">
			<div class="flex items-baseline gap-2">
				<ScheduleTitle
					schedule={flight}
					as="span"
					showRoute={false}
					tooltip={false}
					link={false}
					class="text-lg"
				/>
				<FormattedDatetime value={departureDateTime} class="text-muted-foreground leading-none">
					{formatFlightDate(departureDateTime)}
				</FormattedDatetime>
			</div>
			<div data-route class="flex-1 flex items-start gap-4">
				<div class="flex flex-col">
					<AirportCode airport={flight.originAirport} link={false} class="text-3xl ml-auto" />
					<FormattedDatetime
						value={departureDateTime}
						class="text-sm text-muted-foreground text-right"
					>
						{formatFlightTime(departureDateTime)}
					</FormattedDatetime>
				</div>
				<div class="flex-1 flex-col justify-center items-center relative mt-4">
					<div class="h-[2px] -mt-[1px] w-full bg-muted-foreground/40 flex-1 rounded-full"></div>
					<Plane
						class="size-8 rotate-45 absolute inset-x-1/2 -ml-4 -top-4 fill-foreground"
						strokeWidth="0"
					/>
					<div class="w-full text-center mt-3.5 text-sm text-muted-foreground">
						{formatFlightDuration(departureDateTime, arrivalDateTime)}
					</div>
				</div>
				<div class="flex flex-col">
					<AirportCode airport={flight.destinationAirport} link={false} class="text-3xl" />
					<FormattedDatetime
						value={arrivalDateTime}
						class="text-sm text-muted-foreground text-right"
					>
						{formatFlightTime(arrivalDateTime, { plusMinusDaysFrom: departureDateTime })}
					</FormattedDatetime>
				</div>
			</div>
		</div>
		<div data-info class="flex-0 border-l">
			<DefinitionList.Root
				class="grid grid-cols-[auto_1fr] items-center gap-x-2 gap-y-2 [&>dt]:text-right [&>dt]:tracking-tight [&>dt]:text-xxs [&>dt]:uppercase [&>dd]:!m-0"
			>
				<DefinitionList.Item title="Seat" class="flex items-center gap-2">
					{#snippet titleChild()}<Armchair class="size-5" />{/snippet}
					3A
					<Button variant="secondary" size="xs">Change seat</Button>
				</DefinitionList.Item>
				<DefinitionList.Item title="Aircraft">
					{#snippet titleChild()}<Plane class="size-5" />{/snippet}
					{flight.aircraft?.aircraftType ?? flight.fleet.code}
				</DefinitionList.Item>
				<DefinitionList.Item title="Flight status">
					{#snippet titleChild()}<CircleCheckBig class="size-5" />{/snippet}
					<FlightStatus {flight} />
				</DefinitionList.Item>
			</DefinitionList.Root>
		</div>
	</Card.Content>
</Card.Root>
