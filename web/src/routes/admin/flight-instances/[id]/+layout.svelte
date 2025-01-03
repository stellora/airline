<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import CalendarDays from 'lucide-svelte/icons/calendar-days'
	import Eye from 'lucide-svelte/icons/eye'
	import EyeOff from 'lucide-svelte/icons/eye-off'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'
	import FlightInstanceForm from './flight-instance-form.svelte'
	import Users from 'lucide-svelte/icons/users'

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
			<Dialog.Trigger>
				<Button variant="secondary" size="pageNavbar"><Settings2 /> Edit</Button>
			</Dialog.Trigger>
			<Dialog.Content>
				<Dialog.Header>
					<Dialog.Title>Edit flight instance</Dialog.Title>
				</Dialog.Header>
				<FlightInstanceForm
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
