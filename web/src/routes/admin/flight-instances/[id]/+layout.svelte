<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { buttonVariants } from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import Users from 'lucide-svelte/icons/users'
	import FlightInstanceForm from './flight-instance-form.svelte'

	const { children, data } = $props()
</script>

<PageNav
	tabs={[
		{
			title: 'Overview',
			url: route('/admin/flight-instances/[id]', {
				params: { id: page.params.id },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Passengers',
			url: route('/admin/flight-instances/[id]/passengers', {
				params: { id: page.params.id },
			}),
			icon: Users,
		},
	]}
>
	{#snippet breadcrumbActions()}
		<PageNavbarBreadcrumbActionsDropdownMenu>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/flight-instances/[id]', {
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
		<Dialog.RootByNavigationState id="edit-flight-instance">
			<Dialog.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings2 /> Edit
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>Edit flight instance</Dialog.Title>
				</Dialog.Header>
				<FlightInstanceForm
					flightInstance={data.flightInstance}
					action={route('/admin/flight-instances/[id]', {
						params: { id: page.params.id },
						query: '/update',
					})}
					submitLabel="Save"
					form={data.form}
				/>
			</Dialog.Content>
		</Dialog.RootByNavigationState>
	{/snippet}
</PageNav>

{@render children()}
