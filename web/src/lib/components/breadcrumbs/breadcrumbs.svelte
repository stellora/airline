<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbItem from './breadcrumb-item.svelte'
	import type { BreadcrumbEntry } from './helpers'

	let { entries }: { entries: BreadcrumbEntry[] } = $props()
</script>

<Breadcrumb.Root class="[&>ol>li:not(:last-child)]:max-md:hidden">
	<Breadcrumb.List>
		{#each entries as entry, index (index)}
			<BreadcrumbItem {entry} isLast={index === entries.length - 1} />
		{/each}
	</Breadcrumb.List>
</Breadcrumb.Root>
