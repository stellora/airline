<script lang="ts">
	import { schema } from '$lib/airline.typebox'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Card from '$lib/components/ui/card'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import AircraftForm from '../../../admin/aircraft/aircraft-form.svelte'
	import AircraftList from '../../../admin/aircraft/aircraft-list.svelte'

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
	<Card.Root><AircraftList aircraft={data.aircraft!} showAirline={false} /></Card.Root>
</Page>
