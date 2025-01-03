<script lang="ts">
	import AircraftTypeSelect from '$lib/components/aircraft-type-select.svelte'
	import AirlineSelect from '$lib/components/airline-select.svelte'
	import DateInput from '$lib/components/date-input.svelte'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import { Checkbox } from '$lib/components/ui/checkbox'
	import * as Form from '$lib/components/ui/form'
	import FormFieldGroup from '$lib/components/ui/form/form-field-group.svelte'
	import { Input } from '$lib/components/ui/input'
	import { type FlightInstance } from '$lib/types'
	import { type DateValue, parseDate } from '@internationalized/date'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
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
		flightInstance: Pick<FlightInstance, 'scheduleID'>
		action: string
		submitLabel: string
		form: PageServerData['form']
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
	<FormFieldGroup legend="Flight number">
		<Form.Field {form} name="airline">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Airline</Form.Label>
					<AirlineSelect {...props} bind:value={$formData.airline} {...$constraints.airline} />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="number">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Number</Form.Label>
					<Input
						{...props}
						bind:value={$formData.number}
						autocomplete="off"
						size={8}
						{...$constraints.number}
						class="font-mono w-auto"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</FormFieldGroup>
	<FormFieldGroup legend="Route" horizontal>
		<Form.Field {form} name="originAirport">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Origin airport</Form.Label>
					<Input
						{...props}
						bind:value={$formData.originAirport}
						autocomplete="off"
						size={3}
						{...$constraints.originAirport}
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
						{...$constraints.destinationAirport}
						class="font-mono"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description>IATA code</Form.Description>
			<Form.FieldErrors />
		</Form.Field>
	</FormFieldGroup>
	<Form.Field {form} name="aircraftType">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft type</Form.Label>
				<AircraftTypeSelect
					{...props}
					bind:value={$formData.aircraftType}
					{...$constraints.aircraftType}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>IATA code</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="aircraft">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft</Form.Label>
				{'<AircraftSelect {...props} bind:value={$formData.aircraft} {...$constraints.aircraft} />'}
				<!-- TODO!(sqs) -->
			{/snippet}
		</Form.Control>
		<Form.Description>When known closer to departure</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<FormFieldGroup legend="Timing">
		<Form.Field {form} name="departureDateTime">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Departure</Form.Label>
					<DateInput
						{...props}
						bind:value={() => parseDate($formData.departureDateTime),
						(v: DateValue | undefined) => {
							if (v) {
								$formData.departureDateTime = v.toString()
							}
						}}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</FormFieldGroup>
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
	<Form.Button size="lg">{submitLabel}</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
