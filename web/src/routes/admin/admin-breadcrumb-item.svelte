<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { AdminBreadcrumbEntry } from './admin-breadcrumb'

	let {
		entry,
		isLast
	}: { entry: AdminBreadcrumbEntry | Promise<AdminBreadcrumbEntry>; isLast: boolean } = $props()
</script>

<Breadcrumb.Item>
	{#await entry}
		<Breadcrumb.BreadcrumbEllipsis data-testid="breadcrumb-ellipsis" />
	{:then value}
		<Breadcrumb.Page>{value}</Breadcrumb.Page>
	{/await}
</Breadcrumb.Item>

{#if !isLast}
	<Breadcrumb.Separator role="separator" />
{/if}
