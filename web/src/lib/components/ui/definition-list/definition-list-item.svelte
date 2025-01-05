<script lang="ts">
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		title,
		children,
		...restProps
	}: { title: string; children: Snippet } & HTMLAttributes<HTMLElement> = $props()
</script>

<dt class="font-medium text-sm text-foreground/80">{title}</dt>
<dd {...restProps}>{@render children()}</dd>
