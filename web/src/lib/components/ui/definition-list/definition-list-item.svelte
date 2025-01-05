<script lang="ts">
	import { cn } from '$lib/utils'
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		title,
		children,
		class: className,
		...restProps
	}: { title: string; children: Snippet } & HTMLAttributes<HTMLElement> = $props()
</script>

<dt class="font-medium text-sm text-muted-foreground">{title}</dt>
<dd {...restProps} class={cn('', className)}>{@render children()}</dd>
