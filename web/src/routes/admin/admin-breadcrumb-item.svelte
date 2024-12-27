<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { AdminBreadcrumbEntry } from './admin-breadcrumb'
	import AdminBreadcrumbItem from './admin-breadcrumb-item.svelte'

	let { entry, isLast }: { entry: AdminBreadcrumbEntry; isLast: boolean } = $props()
</script>

{#await entry}
	<Breadcrumb.Item data-testid="breadcrumb-ellipsis">
		<Breadcrumb.BreadcrumbEllipsis />
	</Breadcrumb.Item>
{:then entry}
	{#if !Array.isArray(entry)}
		<Breadcrumb.Item><Breadcrumb.Page>{entry}</Breadcrumb.Page></Breadcrumb.Item>
	{:else}
		{#each entry as entry0, index (index)}
			<AdminBreadcrumbItem entry={entry0} isLast={index === entry.length - 1} />
		{/each}
	{/if}
{/await}

{#if !isLast}
	<Breadcrumb.Separator role="separator" />
{/if}
