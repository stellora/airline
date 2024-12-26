<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import { fade } from 'svelte/transition'
	import AirportForm from './airport-form.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Airports</h1>
	<AirportForm {form} />
	{#if data.airports && data.airports.length > 0}
		<ul class="flex flex-col border rounded-md">
			{#each data.airports as airport (airport.id)}
				<li in:fade out:fade class="border-b last:border-b-0">
					<Button
						variant="link"
						href={`/admin/airports/${airport.id}`}
						class="block p-4 h-[unset] w-full">{airport.title}</Button
					>
				</li>
			{/each}
		</ul>
	{:else}
		<p class="text-muted-foreground">No airports found.</p>
	{/if}
</div>
