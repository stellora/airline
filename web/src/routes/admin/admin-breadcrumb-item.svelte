<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import AdminBreadcrumbItem from './admin-breadcrumb-item.svelte'

	type BreadcrumbEntry = string | Promise<BreadcrumbEntry[]>
	let { item }: { item: string | BreadcrumbEntry[] } = $props()
</script>

{#if typeof item === 'string'}
	<Breadcrumb.Item>
		<Breadcrumb.Page>{item}</Breadcrumb.Page>
	</Breadcrumb.Item>
	<Breadcrumb.Separator />
{:else}
	{#each item as crumb (crumb)}
		{#await crumb}
			<Breadcrumb.BreadcrumbEllipsis />
		{:then crumb}
			<AdminBreadcrumbItem item={crumb} />
		{/await}
	{/each}
{/if}
<!-- TODO!(sqs): this is messy with admin-breadcrumb -->
