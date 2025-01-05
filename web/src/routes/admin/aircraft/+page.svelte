<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import Plus from 'lucide-svelte/icons/plus'
	import AircraftForm from './aircraft-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-aircraft" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New aircraft
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New aircraft</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<AircraftForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/aircraft']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
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
	</Card>
</Page>
