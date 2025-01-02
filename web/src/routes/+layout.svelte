<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import * as Tooltip from '$lib/components/ui/tooltip'
	import '../app.css'

	let { children } = $props()
</script>

<Sidebar.Provider>
	<AppSidebar />
	<div class="w-full">
		<Sidebar.Trigger mobileOnly={true} class="mb-2" />
		<Tooltip.Provider delayDuration={0}>
			{@render children?.()}
		</Tooltip.Provider>
	</div>
</Sidebar.Provider>
