<script lang="ts">
	import { enhance } from '$app/forms'
	import type { ActionData } from './$types'
	let { form }: { form: ActionData } = $props()
</script>

<div class="flex flex-col gap-1">
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
			aria-label="Product title"
			aria-required="true"
		/>
		<button type="submit" class="bg-green-500 text-white" aria-label="Add product">Add</button>
	</form>

	{#if form?.error}
		<p class="text-red-700 p-2" role="alert" aria-live="polite">❌ Error: {form.error}</p>
	{/if}
</div>
