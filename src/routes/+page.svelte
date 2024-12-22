<script>
	import { enhance } from '$app/forms';
	import { fly, slide } from 'svelte/transition';

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
			name="description"
			placeholder="New product description..."
			value={form?.description ?? ''}
			autocomplete="off"
			class="flex-1"
			required
		/>
		<button type="submit" class="bg-green-500 text-white">Add</button>
	</form>

	<ul class="flex flex-col gap-4">
		{#each data.products as product (product.id)}
			<li in:fly={{ y: 20 }} out:slide>
				<form method="POST" action="?/delete" use:enhance class="flex flex-wrap gap-2 items-center">
					<input type="hidden" name="id" value={product.id} />
					<button type="submit">Delete</button>
					<span>{product.description}</span>
				</form>
			</li>
		{/each}
	</ul>
</div>

