<script lang="ts">
	import type { Category, Product } from '$lib/types'
	import Tag from 'lucide-svelte/icons/tag'
	import ProductTitle from './product-title.svelte'

	const { product, categories }: { product: Product; categories?: Category[] } = $props()
</script>

<li class="flex flex-col gap-4 border p-3 rounded-md min-h-32">
	<ProductTitle {product} class="text-lg font-bold leading-none" />
	{#if categories && categories.length > 0}
		<div class="flex gap-2 items-start">
			<Tag size="1.5ch" class="mt-1.5 text-muted-foreground" />
			<ul class="flex gap-2">
				{#each categories as category (category.id)}
					<span class="p-1 text-xs rounded border">{category.title}</span>
				{/each}
			</ul>
		</div>
	{/if}
</li>
