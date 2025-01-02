<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { WithElementRef } from 'bits-ui'
	import type { HTMLAttributes } from 'svelte/elements'

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLTableSectionElement>> = $props()
</script>

<tbody
	bind:this={ref}
	class={cn('[&_tr:last-child]:border-0 [&>tr:hover]:bg-muted/50', className)}
	{...restProps}
>
	{@render children?.()}
</tbody>
