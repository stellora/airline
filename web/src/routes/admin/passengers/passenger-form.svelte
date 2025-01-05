<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	let props: {
		schema:
			| (typeof schema)['/passengers']['POST']['args']['properties']['body']
			| (typeof schema)['/passengers/{id}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/passengers']['POST']['args']['properties']['body']
				| (typeof schema)['/passengers/{id}']['PATCH']['args']['properties']['body']
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
	action={props.action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="passenger-form"
>
	<Form.Field {form} name="name">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Name</Form.Label>
				<Input
					{...props}
					bind:value={$formData.name}
					autocomplete="off"
					class="w-96"
					{...$constraints.name}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>Full legal name</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button>{props.submitLabel}</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
