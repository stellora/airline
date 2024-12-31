<script lang="ts">
	import { page } from '$app/state'
	import { Breadcrumbs } from '$lib/components/breadcrumbs'
	import PageNavbar from '$lib/components/ui/page/page-navbar.svelte'
	import { Separator } from '$lib/components/ui/separator'

	let { children } = $props()
</script>

<PageNavbar>
	<Breadcrumbs entries={page.data.breadcrumbs} />
</PageNavbar>
<Separator class="my-4" />
{@render children?.()}
