<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import Plus from 'lucide-svelte/icons/plus'
	import PassengerForm from './passenger-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-passenger" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New passenger
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New passenger</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<PassengerForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/passengers']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<Page title="Passengers">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[300px]">Name</Table.Head>
					<Table.Head class="text-right"></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.passengers && data.passengers.length > 0}
				<Table.Body>
					{#each data.passengers as passenger (passenger.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell>
								{passenger.name}
							</Table.Cell>
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/admin/passengers/[id]', {
										params: { id: passenger.id.toString() },
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
				<Table.Caption class="mb-4">No passengers found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
