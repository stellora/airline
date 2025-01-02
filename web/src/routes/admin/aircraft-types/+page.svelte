<script lang="ts">
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import { Card } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'

	const { data } = $props()
</script>

<Page title="Aircraft types">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-32">ICAO code</Table.Head>
					<Table.Head>Name</Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.allAircraftTypes && data.allAircraftTypes.length > 0}
				<Table.Body>
					{#each data.allAircraftTypes as aircraftType (aircraftType.icaoCode)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell class="font-bold text-lg">
								<AircraftTypeCode tooltip={false} {aircraftType} />
							</Table.Cell>
							<Table.Cell>{aircraftType.name}</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			{:else}
				<Table.Caption class="mb-4">No aircraft types found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
