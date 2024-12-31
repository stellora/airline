<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import FlightEditForm from './flight-edit-form.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<FlightTitle
			flight={data.flight}
			class="text-4xl font-bold"
			subtitleClass="text-base"
			as="h1"
		/>
	</div>

	<div class="flex flex-wrap-reverse gap-4">
		<Card class="flex-grow-[1]">
			<CardHeader>
				<CardTitle>Flight</CardTitle>
			</CardHeader>
			<FlightEditForm flight={data.flight} {form} />
		</Card>
		<Card class="overflow-hidden flex-grow-[2] basis-[350px] min-w-[350px]">
			<GreatCircleRoute routes={[[data.flight.originAirport, data.flight.destinationAirport]]} />
		</Card>
	</div>

	<Card class="border-destructive self-start">
		<CardContent class="flex gap-4">
			<form method="POST" action="?/setFlightPublished" use:enhance>
				<input type="hidden" name="id" value={data.flight.id} />
				<input type="hidden" name="published" value={data.flight.published ? 'false' : 'true'} />
				<Button type="submit" variant={data.flight.published ? 'outline' : 'default'}>
					{data.flight.published ? 'Unpublish' : 'Publish'}
				</Button>
			</form>
			<form
				method="POST"
				action="?/delete"
				use:enhance={({ cancel }) => {
					if (!confirm('Really delete?')) {
						cancel()
					}
				}}
			>
				<input type="hidden" name="id" value={data.flight.id} />
				<Button type="submit" variant="destructive">Delete flight</Button>
			</form>
		</CardContent>
	</Card>
</div>
