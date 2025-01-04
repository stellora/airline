<script lang="ts">
	import WorldMap from '$lib/components/maps/world-map.svelte'
	import { Card } from '$lib/components/ui/card'
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

<Page title={`${data.airport.iataCode}: ${data.airport.name}`} showTitleHeading>
	{#snippet titleElement()}
		<div class="flex items-baseline gap-2">
			<h1 class="text-4xl font-bold">
				{data.airport.iataCode}
			</h1>
			<span class="text-muted-foreground">{data.airport.name}</span>
		</div>
	{/snippet}

	<Card class="overflow-hidden">
		Timezone: {data.airport.timezoneID}
	</Card>

	<Card class="overflow-hidden">
		<WorldMap features={[geoFeature]} center={geoFeature} />
	</Card>
</Page>
