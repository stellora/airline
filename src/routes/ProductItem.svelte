<script lang="ts">
	import { enhance } from '$app/forms'
	import type { Product } from '$lib/types'
	import { fade } from 'svelte/transition'

	const { product }: { product: Product } = $props()
</script>

<li in:fade out:fade class="flex flex-col gap-4 border border-gray-50/20 p-3">
	<h2 class="text-lg font-bold leading-none">
		{product.title}
		{product.starred ? ' ⭐' : ''}
	</h2>
	<div class="flex flex-wrap gap-2 items-center">
		<form method="POST" action="?/setProductStarred" use:enhance class="flex">
			<input type="hidden" name="id" value={product.id} />
			<input type="hidden" name="starred" value={!product.starred ? 'true' : 'false'} />
			<button type="submit">{product.starred ? 'Unstar' : 'Star'}</button>
		</form>
		<form method="POST" action="?/delete" use:enhance class="flex">
			<input type="hidden" name="id" value={product.id} />
			<button type="submit">Delete</button>
		</form>
	</div>
</li>
