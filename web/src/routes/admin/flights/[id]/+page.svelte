<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<Button variant="outline" href="/admin/flights">← Back</Button>
		<FlightTitle flight={data.flight} class="text-2xl font-bold" as="h1" />
	</div>

	<Card>
		<CardHeader>
			<CardTitle>In airports</CardTitle>
		</CardHeader>
		<CardContent>
			{#if data.flight.airports && data.flight.airports.length > 0}
				<ul class="flex flex-wrap gap-2">
					{#each data.flight.airports as airport (airport.id)}
						<li class="p-2 border text-sm rounded">
							{airport.title}
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-muted-foreground">No airports associated with this flight.</p>
			{/if}
		</CardContent>
	</Card>

	<Card class="border-destructive self-start">
		<CardContent>
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
