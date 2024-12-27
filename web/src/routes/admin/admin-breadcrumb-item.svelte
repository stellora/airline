<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { BreadcrumbItem } from './admin-breadcrumb'

	let { crumb, isLast }: { crumb: BreadcrumbItem | Promise<BreadcrumbItem>; isLast: boolean } =
		$props()
</script>

<Breadcrumb.Item>
	{#await crumb}
		<Breadcrumb.BreadcrumbEllipsis data-testid="breadcrumb-ellipsis" />
	{:then value}
		<Breadcrumb.Page>{value}</Breadcrumb.Page>
	{/await}
</Breadcrumb.Item>

{#if !isLast}
	<Breadcrumb.Separator role="separator" />
{/if}
