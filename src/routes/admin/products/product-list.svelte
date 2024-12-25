<script lang="ts">
	import type { Product } from '$lib/types'
	import ProductItem from './product-item.svelte'

	let { products }: { products: Product[] } = $props()
</script>

<div class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4">
	{#each products as product (product.id)}
		<ProductItem {product} />
	{:else}
		<p class="text-gray-500">No products yet.</p>
	{/each}
</div>
