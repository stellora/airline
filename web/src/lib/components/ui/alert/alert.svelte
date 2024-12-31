<script lang="ts" module>
	import { type VariantProps, tv } from 'tailwind-variants'

	export const alertVariants = tv({
		base: '[&>svg]:text-foreground relative w-full rounded-lg border [&>svg]:absolute',
		variants: {
			variant: {
				default: 'bg-background text-foreground',
				destructive:
					'border-destructive/50 text-destructive dark:border-destructive [&>svg]:text-destructive',
			},
			size: {
				default: 'p-4 [&>svg]:left-4 [&>svg]:top-4 [&>svg~*]:pl-8',
				sm: 'p-2 [&>svg]:left-2 [&>svg]:top-2 [&>svg~*]:pl-6',
			},
		},
		defaultVariants: {
			variant: 'default',
			size: 'default',
		},
	})

	export type AlertVariant = VariantProps<typeof alertVariants>['variant']
	export type AlertSize = VariantProps<typeof alertVariants>['size']
</script>

<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { WithElementRef } from 'bits-ui'
	import type { HTMLAttributes } from 'svelte/elements'

	let {
		ref = $bindable(null),
		class: className,
		variant = 'default',
		size = 'default',
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> & {
		variant?: AlertVariant
		size?: AlertSize
	} = $props()
</script>

<div
	bind:this={ref}
	class={cn(alertVariants({ variant, size }), className)}
	{...restProps}
	role="alert"
>
	{@render children?.()}
</div>
