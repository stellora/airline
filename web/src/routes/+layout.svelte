<script lang="ts">
	import { page } from '$app/state'
	import AppSidebar from '$lib/components/app-sidebar.svelte'
	import PageNavbar from '$lib/components/ui/page/page-navbar.svelte'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import * as Tooltip from '$lib/components/ui/tooltip'
	import '../app.css'

	let { children } = $props()
</script>

<Tooltip.Provider delayDuration={0}>
	<Sidebar.Provider>
		<AppSidebar />
		<div class="w-full overflow-hidden h-screen flex flex-col [&>:first-child]:flex-none">
			<PageNavbar breadcrumbs={page.data.breadcrumbs} />
			<div class="overflow-x-hidden overflow-y-auto">
				<div class="p-4 max-w-screen-lg mx-auto">
					{@render children?.()}
				</div>
			</div>
		</div>
	</Sidebar.Provider>
</Tooltip.Provider>
