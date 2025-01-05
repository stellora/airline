<script lang="ts">
	import WorldMap from '$lib/components/maps/world-map.svelte'
	import * as Card from '$lib/components/ui/card'
	import * as DefinitionList from '$lib/components/ui/definition-list/index.js'
	import Page from '$lib/components/ui/page/page.svelte'
	import type { Feature, Point } from 'geojson'

	let { data } = $props()

	const geoFeature = $derived<Feature<Point>>({
		type: 'Feature',
		properties: { label: data.airport.iataCode },
		geometry: {
			type: 'Point',
			coordinates: [data.airport.point.longitude, data.airport.point.latitude],
		},
	})
</script>

<Page title={`${data.airport.iataCode}: ${data.airport.name}`}>
	<div class="flex flex-wrap-reverse gap-4">
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>
					<h1 class="text-4xl font-bold">
						{data.airport.iataCode}
					</h1>
				</Card.Title>
				<Card.Description>
					{data.airport.name}
				</Card.Description>
			</Card.Header>
			<Card.Content>
				<DefinitionList.Root>
					<DefinitionList.Item title="Location">
						{data.airport.region}, {data.airport.country}
					</DefinitionList.Item>
					<DefinitionList.Item title="Time zone">
						{data.airport.timezoneID}
					</DefinitionList.Item>
				</DefinitionList.Root>
			</Card.Content>
		</Card.Root>

		<Card.Root class="overflow-hidden flex-grow-[1] basis-[250px] min-w-[250px]">
			<WorldMap features={[geoFeature]} center={geoFeature} />
		</Card.Root>
	</div>
</Page>
