<script lang="ts">
	import { enhance } from '$app/forms'
	import { Button } from '$lib/components/ui/button'
	import { Input } from '$lib/components/ui/input'
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
		<Input
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
		<Button type="submit" variant="secondary" aria-label="Add product">Add</Button>
	</form>

	{#if form?.error}
		<p class="text-red-700 p-2" role="alert" aria-live="polite">❌ Error: {form.error}</p>
	{/if}
</div>
