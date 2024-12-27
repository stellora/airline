<script lang="ts">
	import FlightCard from '$lib/components/flight-card.svelte'

	let { data } = $props()
</script>

<div class="flex flex-col gap-10 items-stretch w-full">
	<h1 class="text-4xl font-bold">Book flights</h1>
	<div class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4">
		{#if data.flights}
			{#each data.flights as flight (flight.id)}
				<FlightCard {flight} />
			{:else}
				<p class="text-muted-foreground">No flights yet.</p>
			{/each}
		{/if}
	</div>
</div>
