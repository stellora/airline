<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Separator from '$lib/components/ui/separator/separator.svelte'
	import X from 'lucide-svelte/icons/x'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<Button variant="outline" href="/admin/airports">← Back</Button>
		<h1 class="text-2xl font-bold">{data.airport.title}</h1>
	</div>

	<Card>
		<CardHeader>
			<CardTitle>Flights in airport</CardTitle>
		</CardHeader>
		<CardContent>
			{#if data.flightsInAirport && data.flightsInAirport.length > 0}
				<ul
					class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
					data-testid="flights-in-airport"
				>
					{#each data.flightsInAirport as flight (flight.id)}
						<li class="p-3 border rounded-md flex items-center justify-between gap-2">
							<FlightTitle class="w-full" link {flight} />
							<form
								method="POST"
								action="?/setFlightInAirport"
								use:enhance={({ cancel }) => {
									if (!confirm('Really remove from airport?')) {
										cancel()
									}
								}}
							>
								<input type="hidden" name="airport" value={data.airport.id} />
								<input type="hidden" name="flight" value={flight.id} />
								<input type="hidden" name="value" value="false" />
								<Button
									type="submit"
									variant="ghost"
									size="iconSm"
									aria-label="Remove from airport"
								>
									<X />
								</Button>
							</form>
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-muted-foreground">No flights in this airport.</p>
			{/if}
		</CardContent>
		<Separator />
		<CardContent>
			{#if data.flightsNotInAirport && data.flightsNotInAirport.length > 0}
				<form method="POST" action="?/setFlightInAirport" use:enhance class="flex gap-2">
					<input type="hidden" name="airport" value={data.airport.id} />
					<input type="hidden" name="value" value="true" />
					<select name="flight" class="flex-1">
						{#each data.flightsNotInAirport as flight (flight.id)}
							<option value={flight.id}>{flight.title}</option>
						{/each}
					</select>
					<Button type="submit" variant="secondary">Add flight to airport</Button>
				</form>
			{:else}
				<p class="text-muted-foreground">No flights available to add to this airport.</p>
			{/if}
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
