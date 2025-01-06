<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import AircraftTypeSelect from '$lib/components/aircraft-type-select.svelte'
	import AirlineSelect from '$lib/components/airline-select.svelte'
	import DateRangeInput from '$lib/components/date-range-input.svelte'
	import DaysOfWeekControls from '$lib/components/days-of-week-controls.svelte'
	import FlightDurationInput from '$lib/components/duration-input.svelte'
	import TimeOfDayInput from '$lib/components/time-of-day-input.svelte'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import { Checkbox } from '$lib/components/ui/checkbox'
	import * as Form from '$lib/components/ui/form'
	import FormFieldGroup from '$lib/components/ui/form/form-field-group.svelte'
	import { Input } from '$lib/components/ui/input'
	import { type DaysOfWeek } from '$lib/types'
	import { parseDateTime, type DateValue } from '@internationalized/date'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	const {
		action,
		submitLabel,
		...props
	}: {
		schema:
			| (typeof schema)['/schedules']['POST']['args']['properties']['body']
			| (typeof schema)['/schedules/{id}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/schedules']['POST']['args']['properties']['body']
				| (typeof schema)['/schedules/{id}']['PATCH']['args']['properties']['body']
			>
		>
		action: string
		submitLabel: string
	} = $props()
	const form = superForm(props.data, {
		validators: typebox(props.schema),
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
		dataType: 'json',
	})
	const { form: formData, enhance, message, constraints, errors } = form

	// TODO!(sqs): dedupe with FlightInstanceForm
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="schedule-form"
>
	<FormFieldGroup legend="Flight number" horizontal>
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
	<FormFieldGroup legend="Schedule">
		<Form.Field {form} name="startDate">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Date range</Form.Label>
					<DateRangeInput
						{...props}
						bind:value={() => ({
							start: $formData.startDate ? parseDateTime($formData.startDate) : undefined,
							end: $formData.endDate ? parseDateTime($formData.endDate) : undefined,
						}),
						(v: { start: DateValue | undefined; end: DateValue | undefined }) => {
							if (v.start && v.end) {
								$formData.startDate = v.start.toString()
								$formData.endDate = v.end.toString()
							}
						}}
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Fieldset {form} name="daysOfWeek">
			<Form.Legend>Days of week</Form.Legend>
			<DaysOfWeekControls bind:value={$formData.daysOfWeek as DaysOfWeek} />
		</Form.Fieldset>
		<div class="flex gap-4">
			<Form.Field {form} name="departureTime">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Departure time</Form.Label>
						<TimeOfDayInput
							{...props}
							bind:value={$formData.departureTime}
							autocomplete="off"
							size={10}
							{...$constraints.departureTime}
						/>
					{/snippet}
				</Form.Control>
				<Form.Description
					>Local time at {$errors.originAirport || !$formData.originAirport
						? 'airport'
						: $formData.originAirport}</Form.Description
				>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="durationSec">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Duration</Form.Label>
						<FlightDurationInput
							{...props}
							bind:value={$formData.durationSec}
							autocomplete="off"
							size={10}
							{...$constraints.durationSec}
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>HH:MM format</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
		</div>
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
