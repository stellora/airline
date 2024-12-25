<script lang="ts">
	import { enhance } from '$app/forms'
	import type { ActionData } from './$types'
	let { form }: { form: ActionData } = $props()
</script>

{#if form?.error}
	<p class="text-red-700">{form.error}</p>
{/if}

<form
	method="POST"
	action="?/create"
	use:enhance
	class="flex flex-wrap gap-2"
	data-testid="product-form"
>
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
