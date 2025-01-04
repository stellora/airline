<script lang="ts">
	import AircraftSelect from '$lib/components/aircraft-select.svelte'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Textarea } from '$lib/components/ui/textarea'
	import { type FlightInstance } from '$lib/types'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { LayoutServerData } from './$types'
	import {
		flightInstanceFromManualInputFormSchema,
		flightInstanceFromScheduleFormSchema,
	} from './flight-instance-form'

	const {
		flightInstance,
		action,
		submitLabel,
		...props
	}: {
		flightInstance: Pick<FlightInstance, 'scheduleID' | 'airline'>
		action: string
		submitLabel: string
		form: LayoutServerData['form']
	} = $props()

	const isFromManualInput = flightInstance.scheduleID === undefined

	const form = superForm(props.form, {
		validators: typebox(
			isFromManualInput
				? flightInstanceFromManualInputFormSchema
				: flightInstanceFromScheduleFormSchema,
		),
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
		dataType: 'json',
	})
	const { form: formData, enhance, message, constraints } = form

	// TODO!(sqs): dedupe with FlightScheduleForm
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="flight-instance-form"
>
	<Form.Field {form} name="aircraft">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft</Form.Label>
				<AircraftSelect
					{...props}
					byAirline={flightInstance.airline.id}
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
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
