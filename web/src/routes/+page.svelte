<script lang="ts">
	import FlightCard from '$lib/components/flight-card.svelte'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<div class="p-4">
	<Page title="Book flights">
		<div class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4">
			{#if data.flights}
				{#each data.flights as flight (flight.id)}
					<FlightCard schedule={flight} />
				{:else}
					<p class="text-muted-foreground">No flights found.</p>
				{/each}
			{/if}
		</div>
	</Page>
</div>
