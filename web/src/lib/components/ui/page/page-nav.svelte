<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import type { Snippet } from 'svelte'
	import type { PageNavbarTab } from './page-navbar-tabs.svelte'
	import PageNavbarTabs from './page-navbar-tabs.svelte'
	import { PAGE_NAVBAR_PORTAL_ID } from './page-navbar.svelte'

	const {
		breadcrumbItems,
		breadcrumbActions,
		tabs,
		actions,
	}: {
		breadcrumbItems?: Snippet
		breadcrumbActions?: Snippet
		tabs?: PageNavbarTab[]
		actions?: Snippet
	} = $props()
</script>

{#if breadcrumbItems}
	<Portal target="[aria-label=breadcrumb]">
		{@render breadcrumbItems()}
	</Portal>
{/if}

<Portal target={PAGE_NAVBAR_PORTAL_ID}>
	<div class="flex gap-3">
		{@render breadcrumbActions?.()}

		{#if tabs}
			<PageNavbarTabs {tabs} />
		{/if}

		<div class="flex-1"></div>
		{@render actions?.()}
	</div>
</Portal>
