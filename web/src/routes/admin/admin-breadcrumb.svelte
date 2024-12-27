<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import AdminBreadcrumbItem from './admin-breadcrumb-item.svelte'

	type BreadcrumbEntry = (string | Promise<BreadcrumbEntry[]>)[]
	let { breadcrumbs }: { breadcrumbs: BreadcrumbEntry } = $props()
</script>

<Breadcrumb.Root>
	<Breadcrumb.List>
		{#each breadcrumbs as crumb (crumb)}
			{#await crumb}
				<Breadcrumb.BreadcrumbEllipsis />
			{:then crumbs}
				<AdminBreadcrumbItem item={crumbs} />
			{/await}
		{/each}
	</Breadcrumb.List>
</Breadcrumb.Root>
