<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import { fade } from 'svelte/transition'
	import CategoryForm from './category-form.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Categories</h1>
	<CategoryForm {form} />
	{#if data.categories.length > 0}
		<ul class="flex flex-col border rounded-md">
			{#each data.categories as category (category.id)}
				<li in:fade out:fade class="border-b last:border-b-0">
					<Button
						variant="link"
						href={`/admin/categories/${category.id}`}
						class="block p-4 h-[unset] w-full">{category.title}</Button
					>
				</li>
			{/each}
		</ul>
	{:else}
		<p class="text-muted-foreground">No categories found.</p>
	{/if}
</div>
