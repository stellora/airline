<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import FleetTitle from '$lib/components/fleet-title.svelte'
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import Plus from 'lucide-svelte/icons/plus'
	import FleetForm from './fleet-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-fleet" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New fleet
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New fleet</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FleetForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/airlines/{airlineSpec}/fleets']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<Page title={`Fleets - ${data.airline.iataCode}`}>
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Fleet code</Table.Head>
					<Table.Head>Description</Table.Head>
					<Table.Head class="text-right"></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.fleets && data.fleets.length > 0}
				<Table.Body>
					{#each data.fleets as fleet (fleet.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell class="font-semibold"><FleetTitle {fleet} /></Table.Cell>
							<Table.Cell class="text-muted-foreground">
								{fleet.description}
							</Table.Cell>
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
										params: {
											airlineSpec: fleet.airline.iataCode,
											fleetSpec: fleet.code,
										},
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
				<Table.Caption class="mb-4">No fleets found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
