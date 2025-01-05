<script lang="ts">
	import Distance from '$lib/components/distance.svelte'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import * as Card from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`Scheduled flights - ${data.airline.iataCode}`}>
	<Card.Root>
		<Card.Content>
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
		</Card.Content>
	</Card.Root>
</Page>
