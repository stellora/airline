<script lang="ts">
	import type { Product } from '$lib/types'

	const { product }: { product: Product } = $props()
</script>

<li class="flex flex-col gap-4 border border-gray-50/20 p-3 rounded-md min-h-32">
	<h2 class="text-xl font-bold leading-none">
		{product.title}
		{product.starred ? ' ⭐' : ''}
	</h2>
</li>
