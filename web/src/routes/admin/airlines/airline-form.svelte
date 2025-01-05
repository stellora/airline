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
			| (typeof schema)['/airlines']['POST']['args']['properties']['body']
			| (typeof schema)['/airlines/{airlineSpec}']['PATCH']['args']['properties']['body']
		data: SuperValidated<
			Infer<
				| (typeof schema)['/airlines']['POST']['args']['properties']['body']
				| (typeof schema)['/airlines/{airlineSpec}']['PATCH']['args']['properties']['body']
			>
		>
		action: string
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
	class="flex gap-4 items-center"
	data-testid="airline-form"
>
	<Form.Field {form} name="iataCode">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>IATA code</Form.Label>
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
		<Form.Description>2-letter IATA code</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="name">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Name</Form.Label>
				<Input
					{...props}
					bind:value={$formData.name}
					autocomplete="off"
					class="w-64"
					{...$constraints.name}
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>&nbsp;</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button>Create</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
