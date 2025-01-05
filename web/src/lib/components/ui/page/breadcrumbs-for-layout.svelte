<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import type { Snippet } from 'svelte'
	import { PAGE_BREADCRUMBS_PORTAL_ID } from './page-navbar.svelte'

	const { children, restart = false }: { children: Snippet; restart?: boolean } = $props()
</script>

<Portal target={PAGE_BREADCRUMBS_PORTAL_ID} replace={restart}
	>{@render children()}
	<Breadcrumb.Separator />
</Portal>
