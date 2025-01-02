<script lang="ts">
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import { route } from '$lib/route-helpers'

	let { data } = $props()
</script>

<Page title={`${flightTitle(data.flightSchedule)} schedule`}>
	{#snippet titleElement(className)}
		<FlightTitle flight={data.flightSchedule} class={className} subtitleClass="text-base" as="h1" />
	{/snippet}
	{#snippet titleActions()}
		<Button
			href={route('/admin/flight-schedules/[id]/manage', {
				params: { id: data.flightSchedule.id.toString() },
			})}
			variant="default">Manage flight schedule</Button
		>
	{/snippet}

	<div class="flex flex-wrap-reverse gap-4">
		<Card class="flex-grow-[1]">
			<CardHeader>
				<CardTitle>Route</CardTitle>
			</CardHeader>
		</Card>

		<Card class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute
				routes={[[data.flightSchedule.originAirport, data.flightSchedule.destinationAirport]]}
				detailLevel="high"
			/>
		</Card>
	</div>
</Page>
