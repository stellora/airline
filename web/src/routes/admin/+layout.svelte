<script lang="ts">
	import { page } from '$app/state'
	import PageNavbar from '$lib/components/ui/page/page-navbar.svelte'

	let { children } = $props()
</script>

<PageNavbar breadcrumbs={page.data.breadcrumbs} />
<div class="p-4 max-w-screen-lg mx-auto overflow-x-hidden overflow-y-auto min-h-screen">
	{@render children?.()}
</div>
