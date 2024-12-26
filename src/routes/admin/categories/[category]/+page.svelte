<script lang="ts">
	import { enhance } from '$app/forms'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'
	import Separator from '$lib/components/ui/separator/separator.svelte'

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
				<ul class="flex flex-wrap gap-2">
					{#each data.productsInCategory as product (product.id)}
						<li class="p-2 border rounded-md">{product.title}</li>
					{/each}
				</ul>
			{:else}
				<p class="text-muted-foreground">No products in this category.</p>
			{/if}
		</CardContent>
		<Separator />
		<CardContent>
			<form method="POST" action="?/addProduct" use:enhance class="flex gap-2">
				<input type="hidden" name="category" value={data.category.id} />
				<select name="productId" class="flex-1">
					{#each data.productsNotInCategory as product (product.id)}
						<option value={product.id}>{product.title}</option>
					{/each}
				</select>
				<Button type="submit" variant="secondary">Add product to category</Button>
			</form>
		</CardContent>
	</Card>

	<Card class="border-destructive self-start">
		<CardContent>
			<form method="POST" action="?/delete" use:enhance class="flex">
				<input type="hidden" name="id" value={data.category.id} />
				<Button
					type="submit"
					variant="destructive"
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
		</CardContent>
	</Card>
</div>
