<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import type { Flight } from '$lib/types'
	import { fade } from 'svelte/transition'

	const { flight }: { flight: Flight } = $props()
</script>

<li in:fade out:fade class="flex flex-col gap-4 border p-3 rounded-md">
	<FlightTitle {flight} link class="text-lg font-bold leading-none" />
	<div class="flex flex-wrap gap-2 items-center">
		<form method="POST" action="?/setFlightPublished" use:enhance class="flex">
			<input type="hidden" name="id" value={flight.id} />
			<input type="hidden" name="published" value={!flight.published ? 'true' : 'false'} />
			<Button type="submit" variant="secondary" size="sm"
				>{flight.published ? 'Unstar' : 'Star'}</Button
			>
		</form>
		<Button variant="secondary" size="sm" href={`/admin/flights/${flight.id}`}>Edit</Button>
		<form
			method="POST"
			action="?/delete"
			use:enhance={({ cancel }) => {
				if (!confirm('Really delete?')) {
					cancel()
				}
			}}
		>
			<input type="hidden" name="id" value={flight.id} />
			<Button type="submit" variant="destructive" size="sm">Delete</Button>
		</form>
	</div>
</li>
