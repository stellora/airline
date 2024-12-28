<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent } from '$lib/components/ui/card'

	let { data } = $props()
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

	<p class="text-muted-foreground">{Math.round(data.flight.distanceMiles)} miles</p>

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
