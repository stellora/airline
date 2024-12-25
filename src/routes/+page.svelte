<script lang="ts">
	import ProductCard from '$lib/components/product-card.svelte'

	let { data } = $props()
</script>

<div class="flex flex-col gap-10 items-stretch w-full">
	<h1 class="text-4xl font-bold">Shop</h1>

	<div class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4">
		{#each data.products as product (product.id)}
			<ProductCard {product} />
		{:else}
			<p class="text-gray-500">No products yet.</p>
		{/each}
	</div>
</div>
