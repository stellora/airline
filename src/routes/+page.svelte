<script lang="ts">
	import { enhance } from '$app/forms';
	import { slide } from 'svelte/transition';

	let { data, form } = $props();
</script>

<div class="flex flex-col gap-10 items-stretch w-full max-w-[600px] mx-auto">
	<h1 class="text-4xl font-bold">Products</h1>

	{#if form?.error}
		<p class="text-red-700">{form.error}</p>
	{/if}

	<form method="POST" action="?/create" use:enhance class="flex flex-wrap gap-2">
		<input
			type="text"
			name="title"
			placeholder="New product..."
			value={form?.title ?? ''}
			autocomplete="off"
			class="flex-1"
			required
		/>
		<button type="submit" class="bg-green-500 text-white">Add</button>
	</form>

	<ul class="flex flex-col gap-4">
		{#each data.products as product (product.id)}
			<li in:slide out:slide>
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
</div>

