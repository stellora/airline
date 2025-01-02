<script lang="ts">
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { Card } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightRoute } from '$lib/flight-helpers.js'

	let { data } = $props()
</script>

<Page title={flightRoute(data.route.originAirport, data.route.destinationAirport)}>
	<Card class="max-w-[960px] overflow-hidden">
		<GreatCircleRoute routes={[[data.route.originAirport, data.route.destinationAirport]]} />
	</Card>
</Page>
