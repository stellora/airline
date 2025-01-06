<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import FleetTitle from '$lib/components/fleet-title.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import { cn } from '$lib/utils'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import Trash from 'lucide-svelte/icons/trash'
	import type { ClassNameValue } from 'tailwind-merge'
	import FleetForm from '../fleet-form.svelte'

	let { children, data } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
				params: { airlineSpec: page.params.airlineSpec, fleetSpec: data.fleet.code },
			})}><FleetTitle fleet={data.fleet} tooltip={false} link={false} /></Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav>
	{#snippet breadcrumbActions()}
		<Drawer.DrawerByNavigationState id="edit-fleets" direction="right">
			<PageNavbarBreadcrumbActionsDropdownMenu>
				<DropdownMenu.Group>
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<Drawer.Trigger {...props} class={cn(props.class as ClassNameValue, 'w-full')}>
								<Settings2 /> Edit
							</Drawer.Trigger>
						{/snippet}
					</DropdownMenu.Item>
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<form
								method="POST"
								action={route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
									params: {
										airlineSpec: page.params.airlineSpec,
										fleetSpec: page.params.fleetSpec,
									},
									query: '/delete',
								})}
								use:enhance={({ cancel }) => {
									if (!confirm('Really delete?')) {
										cancel()
									}
								}}
								class="w-full [&>button]:w-full"
							>
								<button type="submit" {...props}>
									<Trash /> Delete...
								</button>
							</form>
						{/snippet}</DropdownMenu.Item
					>
				</DropdownMenu.Group>
			</PageNavbarBreadcrumbActionsDropdownMenu>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit fleet</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<FleetForm
						action={route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
							params: {
								airlineSpec: page.params.airlineSpec,
								fleetSpec: page.params.fleetSpec,
							},
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args'][
							'properties'
						]['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
