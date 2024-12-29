<script lang="ts">
	let { data, form } = $props()
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import { Card, CardContent } from '$lib/components/ui/card'
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-4xl font-bold">Routes</h1>
	<Card>
		<CardContent>
			{#await data.routes}
				<div class="text-muted-foreground">Loading...</div>
			{:then routes}
				{#if routes && routes.length > 0}
					<ul class="grid grid-cols-[repeat(auto-fill,minmax(150px,1fr))] gap-4">
						{#each routes as route (route.originAirport.id + ':' + route.destinationAirport.id)}
							<li class="p-3 border rounded-md flex flex-col items-center justify-between gap-2">
								<span class="text-xl"
									><AirportCode link airport={route.originAirport} />&ndash;<AirportCode
										link
										airport={route.destinationAirport}
									/></span
								>
								<ul class="flex gap-2 justify-around w-full">
									<li class="text-muted-foreground whitespace-nowrap text-sm">
										{route.flightsCount}
										flight{route.flightsCount !== 1 ? 's' : ''}
									</li>
									<li class="text-muted-foreground whitespace-nowrap text-sm">
										<Distance distanceMiles={route.distanceMiles} />
									</li>
								</ul>
							</li>
						{/each}
					</ul>
				{:else}
					<p class="text-muted-foreground">No routes.</p>
				{/if}
			{/await}
		</CardContent>
	</Card>
</div>
