<script lang="ts">
	import { PAGE_NAVBAR_ACTIONS_ID } from '$lib/components/ui/page/page-navbar.svelte'

	import { page } from '$app/state'
	import Portal from '$lib/components/portal.svelte'
	import PageNavbarTabs, {
		type PageNavbarTabsItem,
	} from '$lib/components/ui/page/page-navbar-tabs.svelte'
	import { route } from '$lib/route-helpers'
	import Settings2 from 'lucide-svelte/icons/settings-2'
	import SquareMenu from 'lucide-svelte/icons/square-menu'

	const { children } = $props()

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
	<PageNavbarTabs {tabs} />
</Portal>

{@render children()}
