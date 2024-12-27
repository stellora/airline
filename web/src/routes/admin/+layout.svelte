<script lang="ts">
	import { page } from '$app/state'
	import AdminBreadcrumb from './admin-breadcrumb.svelte'

	let { children } = $props()
</script>

<AdminBreadcrumb entries={page.data.breadcrumbs} />
{@render children?.()}
