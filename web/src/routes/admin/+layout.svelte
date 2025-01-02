<script lang="ts">
	import { page } from '$app/state'
	import { Breadcrumbs } from '$lib/components/breadcrumbs'
	import PageNavbar from '$lib/components/ui/page/page-navbar.svelte'

	let { children } = $props()
</script>

<PageNavbar class="px-4 py-2 sticky top-0 z-50 border-b bg-background h-[47px]">
	<Breadcrumbs entries={page.data.breadcrumbs} />
</PageNavbar>
<div class="p-4">
	{@render children?.()}
</div>
