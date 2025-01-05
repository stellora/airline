<script lang="ts">
	import { enhance } from '$app/forms'
	import { Button } from '$lib/components/ui/button'
	import * as Card from '$lib/components/ui/card/index.js'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`Passenger: ${data.passenger.name}`}>
	<Card.Root>
		<Card.Header
			><Card.Title>{data.passenger.name}</Card.Title>
			<Card.Description>Passenger</Card.Description></Card.Header
		>
		<Card.Content />
	</Card.Root>
</Page>
