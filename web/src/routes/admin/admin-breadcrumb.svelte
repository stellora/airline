<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { AdminBreadcrumbEntry } from './admin-breadcrumb'
	import AdminBreadcrumbItem from './admin-breadcrumb-item.svelte'

	let { breadcrumbs }: { breadcrumbs: AdminBreadcrumbEntry } = $props()
</script>

<Breadcrumb.Root>
	<Breadcrumb.List>
		{#each breadcrumbs as crumb, index (index)}
			<AdminBreadcrumbItem {crumb} isLast={index === breadcrumbs.length - 1} />
		{/each}
	</Breadcrumb.List>
</Breadcrumb.Root>
