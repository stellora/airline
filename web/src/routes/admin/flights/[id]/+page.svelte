<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'

	let { data, form } = $props()
</script>

<Page title={flightTitle(data.flight)}>
	{#snippet titleElement(className)}
		<FlightTitle flight={data.flight} class={className} subtitleClass="text-base" as="h1" />
	{/snippet}
	{#snippet actions()}
		<Button href={`/admin/flight-schedules/${data.flight.id}/manage`} variant="default">Manage flight</Button
		>
	{/snippet}

	<div class="flex flex-wrap-reverse gap-4">
		<Card class="flex-grow-[1]">
			<CardHeader>
				<CardTitle>Flight</CardTitle>
			</CardHeader>
		</Card>
		<Card class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flight.originAirport, data.flight.destinationAirport]]}
				detailLevel="high"
			/>
		</Card>
	</div>
</Page>
