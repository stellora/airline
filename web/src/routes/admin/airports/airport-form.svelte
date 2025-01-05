<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	const {
		action,
		submitLabel,
		...props
	}: {
		schema:
			| (typeof schema)['/airports']['POST']['args']['properties']['body']
			| (typeof schema)['/airports/{airportSpec}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/airports']['POST']['args']['properties']['body']
				| (typeof schema)['/airports/{airportSpec}']['PATCH']['args']['properties']['body']
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
	})
	const { form: formData, enhance, message, constraints } = form
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="airport-form"
>
	<Form.Field {form} name="iataCode">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Airport code</Form.Label>
				<Input
					{...props}
					bind:value={$formData.iataCode}
					autocomplete="off"
					class="font-mono w-32"
					{...$constraints.iataCode}
					oninput={(ev) => {
						ev.currentTarget.value = ev.currentTarget.value.toUpperCase()
					}}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>3-letter IATA code</Form.Description>
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
