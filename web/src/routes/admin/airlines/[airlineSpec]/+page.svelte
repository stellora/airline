<script lang="ts">
	import { enhance } from '$app/forms'
	import Distance from '$lib/components/distance.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`${data.airline.iataCode}: ${data.airline.name}`}>
	{#snippet titleElement()}
		<div class="flex items-baseline gap-2">
			<h1 class="text-4xl font-bold">
				{data.airline.iataCode}
			</h1>
			<span class="text-muted-foreground">{data.airline.name}</span>
		</div>
	{/snippet}

	<Card>
		<CardHeader>
			<CardTitle>Flights on {data.airline.iataCode}</CardTitle>
		</CardHeader>
		<CardContent>
			{#await data.flightSchedules}
				<div class="text-muted-foreground">Loading...</div>
			{:then flightSchedules}
				{#if flightSchedules && flightSchedules.length > 0}
					<ul
						class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
						data-testid="flights-to-from-airline"
					>
						{#each flightSchedules as flight (flight.id)}
							<li class="p-3 border rounded-md flex items-center justify-between gap-2">
								<FlightTitle class="w-full" link {flight} />
								<span class="text-muted-foreground whitespace-nowrap text-sm">
									<Distance distanceMiles={flight.distanceMiles} />
								</span>
							</li>
						{/each}
					</ul>
				{:else}
					<p class="text-muted-foreground">No flights.</p>
				{/if}
			{/await}
		</CardContent>
	</Card>

	{#await data.flightSchedules then flights}
		{#if flights && flights.length > 0}
			<Card class="overflow-hidden">
				<CardHeader>
					<CardTitle>Route map</CardTitle>
				</CardHeader>
				<GreatCircleRoute
					routes={flights.map((flight) => [flight.originAirport, flight.destinationAirport])}
				/>
			</Card>
		{/if}
	{/await}

	<Card class="border-destructive self-start">
		<CardContent>
			<form
				method="POST"
				action="?/delete"
				use:enhance={({ cancel }) => {
					if (!confirm('Really delete?')) {
						cancel()
					}
				}}
			>
				<Button type="submit" variant="destructive">Delete airline</Button>
			</form>
		</CardContent>
	</Card>
</Page>
