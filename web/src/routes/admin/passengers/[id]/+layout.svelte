<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Settings_2 from 'lucide-svelte/icons/settings-2'
	import Trash from 'lucide-svelte/icons/trash'
	import PassengerForm from '../passenger-form.svelte'

	const { data, children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/admin/passengers/[id]', {
				params: { id: data.passenger.id.toString() },
			})}>{data.passenger.name}</Breadcrumb.Link
		>
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
							action={route('/admin/passengers/[id]', {
								params: { id: page.params.id },
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
		<Drawer.DrawerByNavigationState id="edit-airport" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings_2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit airport</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<PassengerForm
						action={route('/admin/passengers/[id]', {
							params: { id: page.params.id },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/passengers/{id}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
