<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import AirlineIcon from '$lib/components/airline-icon.svelte'
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { Card } from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import Plus from 'lucide-svelte/icons/plus'
	import Settings_2 from 'lucide-svelte/icons/settings-2'
	import AirlineForm from './airline-form.svelte'

	let { data } = $props()
</script>

<PageNav>
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="new-airline" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Plus /> New airline
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>New airline</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<AirlineForm
						action="?/create"
						submitLabel="Create"
						data={data.form}
						schema={schema['/airlines']['POST']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

<Page title="Airlines">
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">IATA code</Table.Head>
					<Table.Head>Name</Table.Head>
					<Table.Head></Table.Head>
				</Table.Row>
			</Table.Header>
			{#if data.airlines && data.airlines.length > 0}
				<Table.Body>
					{#each data.airlines as airline (airline.id)}
						<Table.Row class="stretched-link-container group">
							<Table.Cell class="font-bold text-lg">
								<AirlineCode tooltip={false} {airline} />
							</Table.Cell>
							<Table.Cell>
								<div class="flex items-center gap-2">
									<AirlineIcon {airline} />
									{airline.name}
								</div>
							</Table.Cell>
							<Table.Cell>
								<div class="flex items-center justify-end gap-2">
									<Button
										variant="outline"
										href={route('/admin/airlines/[airlineSpec]', {
											params: { airlineSpec: airline.iataCode },
										})}
									>
										<Settings_2 /> Edit public info
									</Button>
									<Button
										variant="secondary"
										href={route('/admin/airlines/[airlineSpec]', {
											params: { airlineSpec: airline.iataCode },
										})}
									>
										Manage...
									</Button>
								</div>
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			{:else}
				<Table.Caption class="mb-4">No airlines found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
