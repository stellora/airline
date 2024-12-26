<script lang="ts">
	import type { Product } from '$lib/types'
	import ProductTitle from './product-title.svelte'

	const { product }: { product: Product } = $props()
</script>

<li class="flex flex-col gap-4 border p-3 rounded-md min-h-32">
	<ProductTitle {product} class="text-lg font-bold leading-none" />
</li>
