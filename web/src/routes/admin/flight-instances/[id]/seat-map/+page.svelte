<script lang="ts">
	import SeatMap from '$lib/components/seat-map/seat-map.svelte'
	import * as Card from '$lib/components/ui/card/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import { formatFlightDate } from '$lib/datetime-helpers'
	import { flightTitle } from '$lib/flight-helpers'
	import { parseZonedDateTime } from '@internationalized/date'

	let { data } = $props()
</script>

<Page
	title={`Passengers on ${flightTitle(data.flightInstance)} ${formatFlightDate(parseZonedDateTime(data.flightInstance.departureDateTime))}`}
>
	<Card.Root class="self-center overflow-hidden">
		<Card.Header>
			<Card.Title>Seat map</Card.Title>
		</Card.Header>
		<Card.CardContent>
			<SeatMap seatAssignments={data.seatAssignments} />
		</Card.CardContent>
	</Card.Root>
</Page>
