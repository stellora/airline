<script lang="ts">
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		title,
		children,
		...restProps
	}: { title: string; children: Snippet } & HTMLAttributes<HTMLElement> = $props()
</script>

<dt class="font-semibold">{title}</dt>
<dd {...restProps}>{@render children()}</dd>
