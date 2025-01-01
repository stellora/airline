<script lang="ts">
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import { formSchema } from './airline-form'

	const props: { form: PageServerData['form']; action: string } = $props()
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
					maxlength={2}
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
				<Input {...props} bind:value={$formData.name} autocomplete="off" class="w-64" />
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
