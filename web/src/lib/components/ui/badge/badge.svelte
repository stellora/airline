<script lang="ts" module>
	import { type VariantProps, tv } from 'tailwind-variants'

	export const badgeVariants = tv({
		base: 'focus:ring-ring inline-flex items-center gap-1.5 rounded-full border font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 [&_svg]:pointer-events-none [&_svg]:shrink-0',
		variants: {
			variant: {
				default:
					'bg-primary text-primary-foreground [&:is(a)]:hover:bg-primary/80 border-transparent',
				secondary:
					'bg-secondary text-secondary-foreground [&:is(a)]:hover:bg-secondary/80 border-transparent',
				destructive:
					'bg-destructive text-destructive-foreground [&:is(a)]:hover:bg-destructive/80 border-transparent',
				outline: 'text-foreground',
			},
			size: {
				default: 'px-2.5 py-0.5 text-xs [&_svg]:size-4',
				xl: 'px-3 py-1.5 text-xl [&_svg]:size-6',
				'4xl': 'px-4 py-2 text-4xl [&_svg]:size-9',
			},
		},
		defaultVariants: {
			variant: 'default',
			size: 'default',
		},
	})

	export type BadgeVariant = VariantProps<typeof badgeVariants>['variant']
</script>

<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { WithElementRef } from 'bits-ui'
	import type { HTMLAnchorAttributes } from 'svelte/elements'
	import type { ClassNameValue } from 'tailwind-merge'

	let {
		ref = $bindable(null),
		href,
		class: className,
		variant = 'default',
		size = 'default',
		children,
		...restProps
	}: WithElementRef<HTMLAnchorAttributes> & {
		variant?: BadgeVariant
		size?: VariantProps<typeof badgeVariants>['size']
	} = $props()
</script>

<svelte:element
	this={href ? 'a' : 'span'}
	bind:this={ref}
	{href}
	class={cn(badgeVariants({ variant, size, className: className as ClassNameValue }))}
	{...restProps}
>
	{@render children?.()}
</svelte:element>
