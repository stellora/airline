<script module>
	export const PAGE_NAVBAR_PORTAL_ID = '#page-navbar-portal'
</script>

<script lang="ts">
	import Breadcrumbs from '$lib/components/breadcrumbs/breadcrumbs.svelte'
	import type { BreadcrumbEntry } from '$lib/components/breadcrumbs/helpers'

	const { breadcrumbs }: { breadcrumbs: BreadcrumbEntry[] } = $props()
</script>

<div class="flex items-center gap-2 px-4 py-2 sticky top-0 z-50 border-b bg-background h-[47px]">
	<Breadcrumbs entries={breadcrumbs} />
	<div id={PAGE_NAVBAR_PORTAL_ID.slice(1)} class="flex-1">
		<!-- Use <Portal target={PAGE_NAVBAR_BREADCRUMB_ACTIONS_ID}> to supply content. -->
	</div>
</div>
