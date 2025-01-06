<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import AircraftSelect from '$lib/components/aircraft-select.svelte'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Textarea } from '$lib/components/ui/textarea'
	import { type Flight } from '$lib/types'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm, type FormPath, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	const {
		flight,
		action,
		submitLabel,
		...props
	}: {
		flight: Pick<Flight, 'scheduleID' | 'airline'>
		schema:
			| (typeof schema)['/flights']['POST']['args']['properties']['body']
			| (typeof schema)['/flights/{id}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/flights']['POST']['args']['properties']['body']
				| (typeof schema)['/flights/{id}']['PATCH']['args']['properties']['body']
			>
		>
		action: string
		submitLabel: string
	} = $props()

	// TODO!(sqs): add more form fields if from manual input
	const isFromManualInput = flight.scheduleID === undefined

	const form = superForm(props.data, {
		validators: typebox(props.schema),
		onSubmit({ jsonData, validators }) {
			if (!$formData.aircraft) {
				$formData.aircraft = undefined
			}

			// Only submit tainted fields.
			const taintedData: Partial<typeof $formData> = Object.fromEntries(
				Object.entries($formData).filter(([key]) => {
					return isTainted(key as FormPath<typeof $formData>)
				}),
			)
			jsonData(taintedData)
		},
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
		dataType: 'json',
	})
	const { form: formData, enhance, message, constraints, errors, isTainted } = form

	// TODO!(sqs): dedupe with ScheduleForm
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="flight-form"
>
	aircraft={JSON.stringify($formData.aircraft)}
	<Form.Field {form} name="aircraft">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft</Form.Label>
				<AircraftSelect
					{...props}
					byAirline={flight.airline.id}
					bind:value={$formData.aircraft}
					{...$constraints.aircraft}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>When known closer to departure</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="notes" class="w-full">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Notes</Form.Label>
				<Textarea
					{...props}
					bind:value={$formData.notes}
					{...$constraints.notes}
					class="h-[130px]"
				/>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button size="lg">{submitLabel}</Form.Button>
	{JSON.stringify($errors)}
	{JSON.stringify($message)}
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
