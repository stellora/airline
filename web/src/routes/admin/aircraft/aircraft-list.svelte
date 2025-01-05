<script lang="ts">
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import { Button } from '$lib/components/ui/button'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import type { Aircraft } from '$lib/types'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'

	const { aircraft, showAirline = true }: { aircraft: Aircraft[]; showAirline?: boolean } = $props()
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			{#if showAirline}<Table.Head class="w-[100px]">Airline</Table.Head>{/if}
			<Table.Head class="w-[130px]">Registration</Table.Head>
			<Table.Head>Type</Table.Head>
			<Table.Head class="text-right"></Table.Head>
		</Table.Row>
	</Table.Header>
	{#if aircraft && aircraft.length > 0}
		<Table.Body>
			{#each aircraft as aircraft (aircraft.id)}
				<Table.Row class="stretched-link-container group">
					{#if showAirline}
						<Table.Cell>
							<AirlineCode airline={aircraft.airline} icon={true} />
						</Table.Cell>
					{/if}
					<Table.Cell class="font-bold font-mono">
						{aircraft.registration}
					</Table.Cell>
					<Table.Cell>
						<AircraftTypeCode aircraftType={aircraft.aircraftType} showName={true} />
					</Table.Cell>
					<Table.Cell class="text-right">
						<Button
							variant="link"
							href={route('/admin/aircraft/[aircraftSpec]', {
								params: { aircraftSpec: aircraft.registration },
							})}
							class="stretched-link h-auto p-1 opacity-35 group-hover:opacity-100"
						>
							<ChevronRight />
						</Button>
					</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	{:else}
		<Table.Caption class="mb-4">No aircraft found</Table.Caption>
	{/if}
</Table.Root>
