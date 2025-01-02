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

<Page title={`${data.airport.iataCode}: ${data.airport.name}`}>
	{#snippet titleElement()}
		<div class="flex items-baseline gap-2">
			<h1 class="text-4xl font-bold">
				{data.airport.iataCode}
			</h1>
			<span class="text-muted-foreground">{data.airport.name}</span>
		</div>
	{/snippet}

	<Card>
		<CardHeader>
			<CardTitle>Flights to/from {data.airport.iataCode}</CardTitle>
		</CardHeader>
		{#await data.flightSchedules then flightSchedules}
			{#if flightSchedules && flightSchedules.length > 0}
				<div class="max-w-[600px]">
					<GreatCircleRoute
						routes={flightSchedules.map((flight) => [
							flight.originAirport,
							flight.destinationAirport,
						])}
					/>
				</div>
			{/if}
		{/await}
		<CardContent>
			{#await data.flightSchedules}
				<div class="text-muted-foreground">Loading...</div>
			{:then flightSchedules}
				{#if flightSchedules && flightSchedules.length > 0}
					<ul
						class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
						data-testid="flights-to-from-airport"
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
				<Button type="submit" variant="destructive">Delete airport</Button>
			</form>
		</CardContent>
	</Card>
</Page>
