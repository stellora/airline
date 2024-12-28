<script lang="ts">
	import { page } from '$app/state'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbItem from './breadcrumb-item.svelte'
	import type { BreadcrumbEntry } from './helpers'
	const { entry, isLast }: { entry: BreadcrumbEntry; isLast: boolean } = $props()
</script>

{#await entry}
	<Breadcrumb.Item data-testid="breadcrumb-ellipsis">
		<Breadcrumb.BreadcrumbEllipsis />
	</Breadcrumb.Item>
{:then entry}
	{#if !Array.isArray(entry)}
		<Breadcrumb.Item>
			{#if typeof entry === 'string'}
				<Breadcrumb.Page>{entry}</Breadcrumb.Page>
			{:else if entry.url === page.url.pathname}
				<Breadcrumb.Page>{entry.title}</Breadcrumb.Page>
			{:else}
				<Breadcrumb.Link href={entry.url}>{entry.title}</Breadcrumb.Link>
			{/if}
		</Breadcrumb.Item>
	{:else}
		{#each entry as entry0, index (index)}
			<BreadcrumbItem entry={entry0} isLast={index === entry.length - 1} />
		{/each}
	{/if}
{/await}

{#if !isLast}
	<Breadcrumb.Separator role="separator" />
{/if}
