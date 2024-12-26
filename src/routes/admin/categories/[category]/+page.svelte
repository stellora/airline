<script lang="ts">
	import { enhance } from '$app/forms'
	import ProductTitle from '$lib/components/product-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Separator from '$lib/components/ui/separator/separator.svelte'
	import X from 'lucide-svelte/icons/x'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<Button variant="outline" href="/admin/categories">← Back</Button>
		<h1 class="text-2xl font-bold">{data.category.title}</h1>
	</div>

	<Card>
		<CardHeader>
			<CardTitle>Products in category</CardTitle>
		</CardHeader>
		<CardContent>
			{#if data.productsInCategory.length > 0}
				<ul
					class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-4"
					data-testid="products-in-category"
				>
					{#each data.productsInCategory as product (product.id)}
						<li class="p-3 border rounded-md flex items-center justify-between gap-2">
							<ProductTitle class="w-full" link {product} />
							<form
								method="POST"
								action="?/setProductInCategory"
								use:enhance={({ cancel }) => {
									if (!confirm('Really remove from category?')) {
										cancel()
									}
								}}
							>
								<input type="hidden" name="category" value={data.category.id} />
								<input type="hidden" name="product" value={product.id} />
								<input type="hidden" name="value" value="false" />
								<Button
									type="submit"
									variant="ghost"
									size="iconSm"
									aria-label="Remove from category"
								>
									<X />
								</Button>
							</form>
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-muted-foreground">No products in this category.</p>
			{/if}
		</CardContent>
		<Separator />
		<CardContent>
			{#if data.productsNotInCategory.length > 0}
				<form method="POST" action="?/setProductInCategory" use:enhance class="flex gap-2">
					<input type="hidden" name="category" value={data.category.id} />
					<input type="hidden" name="value" value="true" />
					<select name="product" class="flex-1">
						{#each data.productsNotInCategory as product (product.id)}
							<option value={product.id}>{product.title}</option>
						{/each}
					</select>
					<Button type="submit" variant="secondary">Add product to category</Button>
				</form>
			{:else}
				<p class="text-muted-foreground">No products available to add to this category.</p>
			{/if}
		</CardContent>
	</Card>

	<Card class="border-destructive self-start">
		<CardContent>
			<form
				method="POST"
				action="?/delete"
				use:enhance={({ cancel }) => {
					if (!confirm('Really delete?')) {
						cancel()
					}
				}}
			>
				<input type="hidden" name="id" value={data.category.id} />
				<Button type="submit" variant="destructive">Delete category</Button>
			</form>
		</CardContent>
	</Card>
</div>
