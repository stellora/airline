<script lang="ts">
	import type { Product } from '$lib/types'
	import { cn } from '$lib/utils'
	import Star from 'lucide-svelte/icons/star'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		product,
		link = false,
		class: className,
		as = 'h2'
	}: {
		product: Product
		link?: boolean
		class: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-center gap-1.5')}>
	{#if link}
		<a href={`/admin/products/${product.id}`}>{product.title}</a>
	{:else}
		{product.title}
	{/if}
	{#if product.starred}
		<Star fill="currentColor" class="inline" size="1.25ch" aria-label="Starred" />
	{/if}
</svelte:element>
