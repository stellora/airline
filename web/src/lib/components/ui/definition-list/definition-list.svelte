<script lang="ts">
	import { cn } from '$lib/utils'
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		children,
		class: className,
		...restProps
	}: { children: Snippet } & HTMLAttributes<HTMLDListElement> = $props()
</script>

<dl {...restProps} class={cn('[&>dd:not(:last-child)]:mb-4', className)}>
	{@render children()}
</dl>
