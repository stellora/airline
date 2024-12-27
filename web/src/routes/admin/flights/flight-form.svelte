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
		data-testid="flight-form"
	>
		<Input
			type="text"
			name="number"
			placeholder="Flight number"
			value={form?.number ?? ''}
			autocomplete="off"
			required
			class="w-32"
			aria-label="Flight number"
			aria-required="true"
		/>
		<div class="flex gap-1 items-center">
			<Input
				type="text"
				name="originAirport"
				placeholder="From"
				value={form?.originAirport ?? ''}
				autocomplete="off"
				required
				maxlength={3}
				class="w-24"
				aria-label="Origin airport IATA code"
				aria-required="true"
			/>
			&ndash;
			<Input
				type="text"
				name="destinationAirport"
				placeholder="To"
				value={form?.destinationAirport ?? ''}
				autocomplete="off"
				required
				maxlength={3}
				class="w-24"
				aria-label="Destination airport IATA code"
				aria-required="true"
			/>
		</div>
		<Button type="submit" variant="secondary" aria-label="Add flight">Add</Button>
	</form>

	{#if form?.error}
		<p class="text-red-700 p-2" role="alert" aria-live="polite">❌ Error: {form.error}</p>
	{/if}
</div>
