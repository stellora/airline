<script lang="ts">
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import AircraftForm from './aircraft-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Dialog.RootByNavigationState id="new-aircraft">
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Plus /> New aircraft</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>New aircraft</Dialog.Title>
				</Dialog.Header>
				<AircraftForm action="?/create" submitLabel="Create" form={data.form} />
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

<Page title="Aircraft">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Airline</Table.Head>
					<Table.Head class="w-[130px]">Registration</Table.Head>
					<Table.Head>Type</Table.Head>
					<Table.Head class="text-right"></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.aircraft && data.aircraft.length > 0}
				<Table.Body>
					{#each data.aircraft as aircraft (aircraft.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell>
								<AirlineCode airline={aircraft.airline} icon={true} />
							</Table.Cell>
							<Table.Cell class="font-bold font-mono">
								{aircraft.registration}
							</Table.Cell>
							<Table.Cell>
								<AircraftTypeCode aircraftType={aircraft.aircraftType} />
							</Table.Cell>
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/admin/aircraft/[aircraftSpec]', {
										params: { aircraftSpec: aircraft.registration },
									})}
									class="stretched-link h-auto p-1 opacity-35 group-hover:opacity-100"
								>
									Manage
								</Button>
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			{:else}
				<Table.Caption class="mb-4">No aircraft found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
