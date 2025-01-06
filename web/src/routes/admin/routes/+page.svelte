<script lang="ts">
	let { data, form } = $props()
	import AirportCode from '$lib/components/airport-code.svelte'
	import Distance from '$lib/components/distance.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route as urlRoute } from '$lib/route-helpers'
</script>

<Page title="Routes">
	<Card>
		<CardContent>
			{#await data.routes}
				<div class="text-muted-foreground">Loading...</div>
			{:then routes}
				{#if routes && routes.length > 0}
					<ul class="grid grid-cols-[repeat(auto-fill,minmax(150px,1fr))] gap-4">
						{#each routes as route (route.originAirport.id + ':' + route.destinationAirport.id)}
							<li
								class="p-3 border rounded-md flex flex-col items-center justify-between gap-1 stretched-link-container"
							>
								<Button
									href={urlRoute('/admin/routes/[route]', {
										params: {
											route: `${route.originAirport.iataCode}-${route.destinationAirport.iataCode}`,
										},
									})}
									variant="link"
									class="text-xl stretched-link gap-0 py-0 h-[unset]"
									><AirportCode airport={route.originAirport} />&ndash;<AirportCode
										airport={route.destinationAirport}
									/></Button
								>
								<ul class="flex gap-3 justify-center w-full">
									<li class="text-muted-foreground whitespace-nowrap text-xs">
										{route.schedulesCount}
										flight{route.schedulesCount !== 1 ? 's' : ''}
									</li>
									<li class="text-muted-foreground whitespace-nowrap text-xs">
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
</Page>
