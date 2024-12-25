<script lang="ts">
	import { enhance } from '$app/forms'
	import ProductTitle from '$lib/components/product-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import type { Product } from '$lib/types'
	import { fade } from 'svelte/transition'

	const { product }: { product: Product } = $props()
</script>

<li in:fade out:fade class="flex flex-col gap-4 border border-gray-50/20 p-3 rounded-md">
	<ProductTitle {product} class="text-lg font-bold leading-none" />
	<div class="flex flex-wrap gap-2 items-center">
		<form method="POST" action="?/setProductStarred" use:enhance class="flex">
			<input type="hidden" name="id" value={product.id} />
			<input type="hidden" name="starred" value={!product.starred ? 'true' : 'false'} />
			<Button type="submit" variant="secondary" size="sm"
				>{product.starred ? 'Unstar' : 'Star'}</Button
			>
		</form>
		<form method="POST" action="?/delete" use:enhance class="flex">
			<input type="hidden" name="id" value={product.id} />
			<Button
				type="submit"
				variant="destructive"
				size="sm"
				onclick={async (event) => {
					event.preventDefault()
					if (confirm('Really delete?')) {
						if (event.currentTarget instanceof HTMLButtonElement) {
							event.currentTarget.form?.submit()
						}
					}
				}}>Delete</Button
			>
		</form>
	</div>
</li>
