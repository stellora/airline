<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { AdminBreadcrumbEntry } from './admin-breadcrumb'
	import AdminBreadcrumbItem from './admin-breadcrumb-item.svelte'

	let { entries }: { entries: AdminBreadcrumbEntry[] } = $props()
</script>

<Breadcrumb.Root>
	<Breadcrumb.List>
		{#each entries as entry, index (index)}
			<AdminBreadcrumbItem {entry} isLast={index === entries.length - 1} />
		{/each}
	</Breadcrumb.List>
</Breadcrumb.Root>
