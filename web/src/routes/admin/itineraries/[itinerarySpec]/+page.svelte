<script lang="ts">
	import ItineraryTitle from '$lib/components/itinerary-title.svelte'
	import * as Card from '$lib/components/ui/card/index.js'
	import * as DefinitionList from '$lib/components/ui/definition-list'
	import Page from '$lib/components/ui/page/page.svelte'
	import ItinerarySegment from './itinerary-segment.svelte'

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
			</DefinitionList.Root>
		</Card.Content>
	</Card.Root>

	{#each data.itinerary.flights as flight (flight.id)}
		<ItinerarySegment {flight} />
	{/each}
</Page>
