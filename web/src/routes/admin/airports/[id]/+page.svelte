<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<Button variant="outline" href="/admin/airports">← Back</Button>
		<h1 class="text-2xl font-bold">{data.airport.iataCode}</h1>
		<p class="text-muted-foreground">{data.airport.name}</p>
	</div>

	<Card>
		<CardHeader>
			<CardTitle>Flights to/from {data.airport.iataCode}</CardTitle>
		</CardHeader>
		<CardContent>
			{#await data.flights}
				<div class="text-muted-foreground">Loading...</div>
			{:then flights}
				{#if flights && flights.length > 0}
					<ul
						class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
						data-testid="flights-to-from-airport"
					>
						{#each flights as flight (flight.id)}
							<li class="p-3 border rounded-md flex items-center justify-between gap-2">
								<FlightTitle class="w-full" link {flight} />
								<span class="text-muted-foreground whitespace-nowrap text-sm">
									{Math.round(flight.distanceMiles)}
									miles
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
				<input type="hidden" name="id" value={data.airport.id} />
				<Button type="submit" variant="destructive">Delete airport</Button>
			</form>
		</CardContent>
	</Card>
</div>
