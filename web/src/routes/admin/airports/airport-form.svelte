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
		data-testid="airport-form"
	>
		<Input
			type="text"
			name="iataCode"
			placeholder="IATA code"
			value={form?.iataCode ?? ''}
			autocomplete="off"
			class="w-24"
			maxlength={3}
			required
			aria-label="Airport IATA code"
			aria-required="true"
			pattern={`[A-Z]{3,3}`}
			oninput={(ev) => {
				ev.currentTarget.value = ev.currentTarget.value.toUpperCase()
			}}
		/>
		<Button type="submit" variant="secondary" aria-label="Add airport">Add</Button>
	</form>

	{#if form?.error}
		<p class="text-red-700 p-2" role="alert" aria-live="polite">❌ Error: {form.error}</p>
	{/if}
</div>
