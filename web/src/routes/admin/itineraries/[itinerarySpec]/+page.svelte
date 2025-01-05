<script lang="ts">
	import ItineraryTitle from '$lib/components/itinerary-title.svelte'
	import * as Card from '$lib/components/ui/card/index.js'
	import * as DefinitionList from '$lib/components/ui/definition-list'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightInstanceTitle } from '$lib/flight-helpers.js'
	import { route } from '$lib/route-helpers.js'

	let { data } = $props()
</script>

<Page title={`${data.itinerary.recordID} (itinerary)`}>
	<Card.Root>
		<Card.Header
			><Card.Title><ItineraryTitle itinerary={data.itinerary} /></Card.Title>
			<Card.Description>Itinerary</Card.Description></Card.Header
		>
		<Card.Content>
			<DefinitionList.Root>
				<DefinitionList.Item title="Passengers">
					<ul>
						{#each data.itinerary.passengers as passenger (passenger.id)}
							<li>{passenger.name}</li>
						{/each}
					</ul>
				</DefinitionList.Item>
				<DefinitionList.Item title="Flights">
					<ul>
						{#each data.itinerary.flights as flight (flight.id)}
							<li>
								<a
									href={route('/admin/flight-instances/[id]', {
										params: { id: flight.id.toString() },
									})}
								>
									{flightInstanceTitle(flight)}
								</a>
							</li>
						{/each}
					</ul>
				</DefinitionList.Item>
			</DefinitionList.Root>
		</Card.Content>
	</Card.Root>
</Page>
