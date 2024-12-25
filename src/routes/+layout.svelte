<script lang="ts">
	import AppSidebar from '$lib/components/app-sidebar.svelte'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import '../app.css'

	let { children } = $props()
</script>

<Sidebar.Provider>
	<AppSidebar />
	<Sidebar.Inset class="p-4">
		<Sidebar.Trigger mobileOnly={true} />
		{@render children?.()}
	</Sidebar.Inset>
</Sidebar.Provider>
