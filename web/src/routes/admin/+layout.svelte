<script lang="ts">
	import { page } from '$app/state'
	import { Breadcrumbs } from '$lib/components/breadcrumbs'

	let { children } = $props()
</script>

<Breadcrumbs entries={page.data.breadcrumbs} />
{@render children?.()}
