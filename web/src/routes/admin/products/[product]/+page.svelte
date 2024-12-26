<script lang="ts">
	import { enhance } from '$app/forms'
	import ProductTitle from '$lib/components/product-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'

	let { data } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<div class="flex items-center gap-4">
		<Button variant="outline" href="/admin/products">← Back</Button>
		<ProductTitle product={data.product} class="text-2xl font-bold" as="h1" />
	</div>

	<Card>
		<CardHeader>
			<CardTitle>In categories</CardTitle>
		</CardHeader>
		<CardContent>
			{#if data.product.categories && data.product.categories.length > 0}
				<ul class="flex flex-wrap gap-2">
					{#each data.product.categories as category (category.id)}
						<li class="p-2 border text-sm rounded">
							{category.title}
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-muted-foreground">No categories associated with this product.</p>
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
				<input type="hidden" name="id" value={data.product.id} />
				<Button type="submit" variant="destructive">Delete product</Button>
			</form>
		</CardContent>
	</Card>
</div>
