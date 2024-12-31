<script lang="ts">
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import { formSchema } from './flight-form'

	const props: { form: PageServerData['form'] } = $props()
	$inspect(props.form)
	const form = superForm(props.form, { validators: typebox(formSchema) })
	const { form: formData, enhance } = form
</script>

<div class="flex flex-col gap-1">
	<form
		method="POST"
		action="?/create"
		use:enhance
		class="flex flex-wrap gap-2"
		data-testid="flight-form"
	>
		<Form.Field {form} name="number">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Flight number</Form.Label>
					<Input {...props} bind:value={$formData.number} autocomplete="off" />
				{/snippet}
			</Form.Control>
			<Form.Description>Airline (2-letter IATA code) + number</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		<fieldset>
			<legend>Hello</legend>
			<Form.Field {form} name="originAirport">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Origin airport</Form.Label>
						<Input
							{...props}
							bind:value={$formData.originAirport}
							autocomplete="off"
							size={3}
							class="font-mono"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>IATA code</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="destinationAirport">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Destination airport</Form.Label>
						<Input
							{...props}
							bind:value={$formData.destinationAirport}
							autocomplete="off"
							size={3}
							class="font-mono"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>IATA code</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
		</fieldset>
		<Form.Button type="submit" variant="secondary" aria-label="Add flight">Add</Form.Button>
	</form>

	{#if form?.error}
		<p class="text-red-700 p-2" role="alert" aria-live="polite">❌ Error: {form.error}</p>
	{/if}
</div>
