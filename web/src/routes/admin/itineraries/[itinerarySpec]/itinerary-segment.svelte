<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js'
	import { flightTitle } from '$lib/flight-helpers'
	import type { Flight } from '$lib/types'

	const { flight }: { flight: Flight } = $props()
</script>

<Card.Root>
	<Card.Content>{flightTitle(flight)}</Card.Content>
</Card.Root>
