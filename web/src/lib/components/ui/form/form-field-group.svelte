<script lang="ts" generics="T extends Record<string, unknown>, U extends _FormPath<T>">
	import { cn } from '$lib/utils.js'
	import type { WithElementRef } from 'bits-ui'
	import type { HTMLFieldsetAttributes } from 'svelte/elements'

	let {
		ref = $bindable(null),
		class: className,
		legend,
		children,
		...restProps
	}: { legend: string } & WithElementRef<HTMLFieldsetAttributes> = $props()
</script>

<fieldset
	bind:this={ref}
	class={cn('rounded-lg border px-4 pt-2 pb-0 w-fit', className)}
	{...restProps}
>
	<legend class="text-base font-medium">{legend}</legend>
	<div class={cn('flex gap-4')}>
		{@render children?.()}
	</div>
</fieldset>
