<script lang="ts">
	import FleetTitle from '$lib/components/fleet-title.svelte'
	import * as Card from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import AircraftList from '../../aircraft/aircraft-list.svelte'

	let { data } = $props()
</script>

<Page title={`Fleet ${data.fleet.code}`}>
	<Card.Root>
		<Card.Header>
			<Card.Title>
				<FleetTitle fleet={data.fleet} link={false} tooltip={false} class="text-4xl" />
			</Card.Title>
			<Card.Description>
				{data.fleet.description}
			</Card.Description>
		</Card.Header>
		<Card.Content />
	</Card.Root>

	<Card.Root>
		<AircraftList aircraft={data.fleetAircraft} showAirline={false} />
	</Card.Root>
</Page>
