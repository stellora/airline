<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import Portal from '$lib/components/portal.svelte'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import PageNavbarOtherActions from '$lib/components/ui/page/page-navbar-other-actions.svelte'
	import PageNavbarTabs, {
		type PageNavbarTabsItem,
	} from '$lib/components/ui/page/page-navbar-tabs.svelte'
	import { PAGE_NAVBAR_ACTIONS_ID } from '$lib/components/ui/page/page-navbar.svelte'
	import { route } from '$lib/route-helpers'
	import Eye from 'lucide-svelte/icons/eye'
	import EyeOff from 'lucide-svelte/icons/eye-off'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'
	import Trash from 'lucide-svelte/icons/trash'

	const { children, data } = $props()

	const tabs: PageNavbarTabsItem[] = [
		{
			title: 'Overview',
			url: route('/admin/flight-schedules/[id]', {
				params: { id: page.params.id },
			}),
			icon: SquareMenu,
		},
		{
			title: 'Manage',
			url: route('/admin/flight-schedules/[id]/manage', {
				params: { id: page.params.id },
			}),
			icon: Settings2,
		},
	]
</script>

<Portal target={PAGE_NAVBAR_ACTIONS_ID}>
	<div class="flex flex-wrap gap-3">
		<PageNavbarOtherActions>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/flight-schedules/[id]', {
								params: { id: page.params.id },
								query: '/setFlightSchedulePublished',
							})}
							use:enhance
							class="w-full [&>button]:w-full"
						>
							<input
								type="hidden"
								name="published"
								value={data.flightSchedule.published ? 'false' : 'true'}
							/>
							<button type="submit" {...props}>
								{#if data.flightSchedule.published}
									<EyeOff /> Unpublish
								{:else}
									<Eye /> Publish
								{/if}
							</button>
						</form>
					{/snippet}
				</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/flight-schedules/[id]', {
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
							<input
								type="hidden"
								name="published"
								value={data.flightSchedule.published ? 'false' : 'true'}
							/>
							<button type="submit" {...props}>
								<Trash /> Delete...
							</button>
						</form>
					{/snippet}</DropdownMenu.Item
				>
			</DropdownMenu.Group></PageNavbarOtherActions
		>
		<PageNavbarTabs {tabs} />
	</div>
</Portal>

{@render children()}
