<script lang="ts">
	import { enhance } from '$app/forms'
	import type { Product } from '$lib/types'
	import { fade } from 'svelte/transition'
	let { products }: { products: Product[] } = $props()
</script>

<ul class="flex flex-col gap-4">
	{#each products as product (product.id)}
		<li in:fade out:fade>
			<form method="POST" action="?/delete" use:enhance class="flex flex-wrap gap-2 items-center">
				<input type="hidden" name="id" value={product.id} />
				<button type="submit">Delete</button>
				<span>{product.title}</span>
			</form>
		</li>
	{:else}
		<p class="text-gray-500">No products yet.</p>
	{/each}
</ul>
