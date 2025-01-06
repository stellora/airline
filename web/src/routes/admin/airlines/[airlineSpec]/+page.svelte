<script lang="ts">
	import AirlineCode from '$lib/components/airline-code.svelte'
	import * as Card from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`${data.airline.iataCode}: ${data.airline.name}`}>
	<Card.Root>
		<Card.Header>
			<Card.Title><AirlineCode airline={data.airline} icon /></Card.Title>
			<Card.Description>
				{data.airline.name}
			</Card.Description>
		</Card.Header>
		<Card.Content />
	</Card.Root>
</Page>
