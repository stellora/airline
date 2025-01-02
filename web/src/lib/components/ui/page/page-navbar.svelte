<script module>
	export const PAGE_NAVBAR_ACTIONS_ID = '#page-navbar-actions'
</script>

<script lang="ts">
	import { cn } from '$lib/utils'
	import type { Snippet } from 'svelte'

	const { children, class: className }: { children?: Snippet; class?: string } = $props()
</script>

<div class={cn('flex flex-wrap items-center gap-2', className)}>
	{@render children?.()}
	<div id={PAGE_NAVBAR_ACTIONS_ID.slice(1)}>
		<!-- Use <Portal target={PAGE_NAVBAR_ACTIONS_ID} to supply content. -->
	</div>
</div>
