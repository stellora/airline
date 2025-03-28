<script lang="ts">
	import FlightCard from '$lib/components/flight-card.svelte'
import Page from '$lib/components/ui/page/page.svelte'
import { isFeatureFlagEnabled } from '$lib/feature-flags'

	let { data } = $props()
</script>

<div class="p-4">
	<Page title="Book flights">
		{#if isFeatureFlagEnabled('payment.apple-pay')}
			<!-- Apple Pay integration would appear here -->
			<!-- <div class="mb-4 p-2 border rounded flex items-center gap-2 bg-gray-50">
				<img src="/apple-pay-logo.svg" alt="Apple Pay" width="60" />
				<span>Book flights quickly with Apple Pay</span>
			</div> -->
		{/if}
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
