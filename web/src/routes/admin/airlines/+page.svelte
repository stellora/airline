<script lang="ts">
	import AirlineCode from '$lib/components/airline-code.svelte'
	import AirlineIcon from '$lib/components/airline-icon.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import * as Table from '$lib/components/ui/table'
	import { route } from '$lib/route-helpers'
	import ChevronRight from 'lucide-svelte/icons/chevron-right'
	import AirlineForm from './airline-form.svelte'

	let { data } = $props()
</script>

<Page title="Airlines">
	<Card class="self-start">
		<CardHeader>
			<CardTitle>New airline</CardTitle>
		</CardHeader>

		<CardContent>
			<AirlineForm form={data.form} action="?/create" />
		</CardContent>
	</Card>
	<Card>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">IATA code</Table.Head>
					<Table.Head>Name</Table.Head>
					<Table.Head class="text-right"></Table.Head>
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
							<Table.Cell class="text-right">
								<Button
									variant="link"
									href={route('/admin/airlines/[airlineSpec]', {
										params: { airlineSpec: airline.iataCode },
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
				<Table.Caption class="mb-4">No airlines found</Table.Caption>
			{/if}
		</Table.Root>
	</Card>
</Page>
