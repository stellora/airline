<script module>
	export const PAGE_NAVBAR_ACTIONS_ID = '#page-navbar-actions'
</script>

<script lang="ts">
	import type { Snippet } from 'svelte'

	const { children }: { children?: Snippet } = $props()
</script>

<div class="flex flex-wrap items-center justify-between gap-4">
	{@render children?.()}
	<div id={PAGE_NAVBAR_ACTIONS_ID.slice(1)}>
		<!-- Use <Portal target={PAGE_NAVBAR_ACTIONS_ID} to supply content. -->
	</div>
</div>
