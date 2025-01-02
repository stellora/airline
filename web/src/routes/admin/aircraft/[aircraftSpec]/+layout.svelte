<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import Trash from 'lucide-svelte/icons/trash'
	import AircraftForm from '../aircraft-form.svelte'

	const { children, data } = $props()
</script>

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
		<Dialog.RootByNavigationState id="edit-aircraft">
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Settings2 /> Edit</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>Edit aircraft</Dialog.Title>
				</Dialog.Header>
				<AircraftForm
					action={route('/admin/aircraft/[aircraftSpec]', {
						params: { aircraftSpec: page.params.aircraftSpec },
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
