<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import type { Snippet } from 'svelte'
	import type { PageNavbarTab } from './page-navbar-tabs.svelte'
	import PageNavbarTabs from './page-navbar-tabs.svelte'
	import { PAGE_NAVBAR_PORTAL_ID } from './page-navbar.svelte'

	const {
		breadcrumbActions,
		tabs,
		actions,
	}: {
		breadcrumbActions?: Snippet
		tabs?: PageNavbarTab[]
		actions?: Snippet
	} = $props()
</script>

<Portal target={PAGE_NAVBAR_PORTAL_ID}>
	{@render breadcrumbActions?.()}

	{#if tabs}
		<PageNavbarTabs {tabs} />
	{/if}

	{#if actions}
		<div class="ml-auto flex">
			{@render actions?.()}
		</div>
	{/if}
</Portal>
