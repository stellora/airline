<script lang="ts">
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import { Checkbox } from '$lib/components/ui/checkbox'
	import * as Form from '$lib/components/ui/form'
	import FormFieldGroup from '$lib/components/ui/form/form-field-group.svelte'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import FlightDateRangeInput from './flight-date-range-input.svelte'
	import { formSchema } from './new-flight-form'

	const props: { form: PageServerData['form'] } = $props()
	const form = superForm(props.form, {
		validators: typebox(formSchema),
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
	})
	const { form: formData, enhance, message } = form
</script>

<form
	method="POST"
	action="?/create"
	use:enhance
	class="flex flex-col gap-6 items-start"
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
		<span class="relative top-[31px] left-[-4px] w-[1px]">&ndash;</span>
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
	<Form.Field {form} name="dateRange">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Date range</Form.Label>
				<FlightDateRangeInput {...props} bind:value={$formData.dateRange} />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<FormFieldGroup legend="Options">
		<Form.Field {form} name="published">
			<Form.Control>
				{#snippet children({ props })}
					<div class="flex items-center gap-2">
						<Checkbox {...props} bind:checked={$formData.published} />
						<Form.Label>Published</Form.Label>
					</div>
				{/snippet}
			</Form.Control>
			<Form.Description>Published flights are immediately available for booking</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</FormFieldGroup>
	<Form.Button size="lg">Create</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
