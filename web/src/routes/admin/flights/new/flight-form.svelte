<script lang="ts">
	import * as Form from '$lib/components/ui/form'
	import FormFieldGroup from '$lib/components/ui/form/form-field-group.svelte'
	import { Input } from '$lib/components/ui/input'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import { formSchema } from './flight-form'

	const props: { form: PageServerData['form'] } = $props()
	const form = superForm(props.form, { validators: typebox(formSchema) })
	const { form: formData, enhance } = form
</script>

<div class="flex flex-col gap-1">
	<form
		method="POST"
		action="?/create"
		use:enhance
		class="flex flex-col gap-4"
		data-testid="flight-form"
	>
		<Form.Field {form} name="number">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Flight number</Form.Label>
					<Input
						{...props}
						bind:value={$formData.number}
						autocomplete="off"
						size={16}
						class="font-mono w-auto"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description>Airline (2-letter IATA code) + number</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
		<FormFieldGroup legend="Route">
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
			<span class="relative top-[38px] left-[-4px] w-[1px]">&ndash;</span>
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
		</FormFieldGroup>
		<Form.Button>Create flight</Form.Button>
	</form>
</div>
