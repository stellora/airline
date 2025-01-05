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
	<div class="flex gap-3">
		{@render breadcrumbActions?.()}

		{#if tabs}
			<PageNavbarTabs {tabs} />
		{/if}

		<div class="flex-1"></div>
		{@render actions?.()}
	</div>
</Portal>
