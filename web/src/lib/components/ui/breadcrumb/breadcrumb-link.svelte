<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { WithElementRef } from 'bits-ui'
	import type { Snippet } from 'svelte'
	import type { HTMLAnchorAttributes } from 'svelte/elements'

	let {
		ref = $bindable(null),
		class: className,
		href = undefined,
		isActive,
		child,
		children,
		...restProps
	}: WithElementRef<HTMLAnchorAttributes> & {
		isActive?: boolean
		child?: Snippet<[{ props: HTMLAnchorAttributes }]>
	} = $props()

	const attrs = $derived({
		class: cn(
			'hover:text-foreground transition-colors',
			{ 'text-foreground': isActive },
			className,
		),
		href,
		...restProps,
	})
</script>

{#if child}
	{@render child({ props: attrs })}
{:else}
	<a bind:this={ref} {...attrs}>
		{@render children?.()}
	</a>
{/if}
