<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { HTMLAttributes } from 'svelte/elements'

	let { class: className, children, ...restProps }: HTMLAttributes<HTMLDivElement> = $props()
</script>

<div class={cn('p-4 overflow-y-auto', className)} {...restProps}>
	{@render children?.()}
</div>
