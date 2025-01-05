<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import AircraftRegistration from '$lib/components/aircraft-registration.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import Trash from 'lucide-svelte/icons/trash'
	import AircraftForm from '../aircraft-form.svelte'

	const { children, data } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/admin/aircraft/[aircraftSpec]', {
				params: { aircraftSpec: data.aircraft.registration },
			})}
			><AircraftRegistration aircraft={data.aircraft} showAircraftType={true} />
		</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav>
	{#snippet breadcrumbActions()}
		<PageNavbarBreadcrumbActionsDropdownMenu>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/aircraft/[aircraftSpec]', {
								params: { aircraftSpec: page.params.aircraftSpec },
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
	{/snippet}
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="edit-aircraft" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit aircraft</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<AircraftForm
						action={route('/admin/aircraft/[aircraftSpec]', {
							params: { aircraftSpec: page.params.aircraftSpec },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/aircraft/{aircraftSpec}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
