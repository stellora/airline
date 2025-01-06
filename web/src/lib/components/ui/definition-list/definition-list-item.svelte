<script lang="ts">
	import { cn } from '$lib/utils'
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		title,
		titleChild,
		children,
		class: className,
		...restProps
	}: {
		title: string
		titleChild?: Snippet
		children: Snippet
	} & HTMLAttributes<HTMLElement> = $props()
</script>

<dt class="font-medium text-sm text-foreground/75">
	{#if titleChild}
		{@render titleChild()}
	{:else}
		{title}
	{/if}
</dt>
<dd {...restProps} class={cn('', className)}>{@render children()}</dd>
