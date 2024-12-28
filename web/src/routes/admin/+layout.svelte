<script lang="ts">
	import { page } from '$app/state'
	import { Breadcrumbs } from '$lib/components/breadcrumbs'
	import { Separator } from '$lib/components/ui/separator'

	let { children } = $props()
</script>

<Breadcrumbs entries={page.data.breadcrumbs} />
<Separator class="my-4" />
{@render children?.()}
