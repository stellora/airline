<script lang="ts">
	import AirlineSelect from '$lib/components/airline-select.svelte'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import { formSchema } from './aircraft-form'

	const {
		action,
		submitLabel,
		...props
	}: { action: string; submitLabel: string; form: PageServerData['form'] } = $props()
	const form = superForm(props.form, {
		validators: typebox(formSchema),
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
	})
	const { form: formData, enhance, message, constraints } = form
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="aircraft-form"
>
	<Form.Field {form} name="registration">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft code</Form.Label>
				<Input
					{...props}
					bind:value={$formData.registration}
					autocomplete="off"
					class="font-mono w-32"
					{...$constraints.registration}
					oninput={(ev) => {
						ev.currentTarget.value = ev.currentTarget.value.toUpperCase()
					}}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>ICAO unique registration</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="aircraftType">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Aircraft type</Form.Label>
				<Input
					{...props}
					bind:value={$formData.aircraftType}
					autocomplete="off"
					class="w-32"
					{...$constraints.aircraftType}
					oninput={(ev) => {
						ev.currentTarget.value = ev.currentTarget.value.toUpperCase()
					}}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>ICAO code</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="airline">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Airline</Form.Label>
				<AirlineSelect {...props} bind:value={$formData.airline} {...$constraints.airline} />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button>{submitLabel}</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
