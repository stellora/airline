<script lang="ts">
	import Distance from '$lib/components/distance.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'

	let { data } = $props()
</script>

<Card>
	<CardHeader>
		<CardTitle>Flights to/from {data.airport.iataCode}</CardTitle>
	</CardHeader>
	{#await data.schedules then schedules}
		{#if schedules && schedules.length > 0}
			<GreatCircleRoute
				routes={schedules.map((flight) => [flight.originAirport, flight.destinationAirport])}
			/>
		{/if}
	{/await}
	<CardContent>
		{#await data.schedules}
			<div class="text-muted-foreground">Loading...</div>
		{:then schedules}
			{#if schedules && schedules.length > 0}
				<ul
					class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
					data-testid="flights-to-from-airport"
				>
					{#each schedules as flight (flight.id)}
						<li
							class="p-3 border rounded-md flex items-center justify-between gap-2 stretched-link-container"
						>
							<ScheduleTitle class="w-full" linkClass="stretched-link" link schedule={flight} />
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
