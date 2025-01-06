<script lang="ts">
	import { enhance } from '$app/forms'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import * as Card from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`${data.airline.iataCode}: ${data.airline.name}`}>
	<Card.Root>
		<Card.Header>
			<Card.Title><AirlineCode airline={data.airline} icon /></Card.Title>
			<Card.Description>
				{data.airline.name}
			</Card.Description>
		</Card.Header>
		<Card.Content />
	</Card.Root>

	{#await data.flightSchedules then flights}
		{#if flights && flights.length > 0}
			<Card.Root class="overflow-hidden">
				<Card.Header>
					<Card.Title>Route map</Card.Title>
				</Card.Header>
				<GreatCircleRoute
					routes={flights.map((flight) => [flight.originAirport, flight.destinationAirport])}
				/>
			</Card.Root>
		{/if}
	{/await}
</Page>
