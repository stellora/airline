<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import { Textarea } from '$lib/components/ui/textarea'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	const {
		action,
		submitLabel,
		...props
	}: {
		schema:
			| (typeof schema)['/airlines/{airlineSpec}/fleets']['POST']['args']['properties']['body']
			| (typeof schema)['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/airlines/{airlineSpec}/fleets']['POST']['args']['properties']['body']
				| (typeof schema)['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args']['properties']['body']
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
	const { form: formData, enhance, message, constraints } = form
</script>

<form
	method="POST"
	{action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="fleet-form"
>
	<Form.Field {form} name="code">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Fleet code</Form.Label>
				<Input
					{...props}
					bind:value={$formData.code}
					autocomplete="off"
					class="font-mono w-64"
					{...$constraints.code}
					oninput={(ev) => {
						ev.currentTarget.value = ev.currentTarget.value.toUpperCase()
					}}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>Unique identifier for the fleet</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="description" class="self-stretch">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Description</Form.Label>
				<Textarea
					{...props}
					bind:value={$formData.description}
					autocomplete="off"
					class="w-full h-24"
					{...$constraints.description}
				/>
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
