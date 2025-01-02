<script lang="ts">
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { superForm } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'
	import type { PageServerData } from './$types'
	import { formSchema } from './airport-form'

	const props: { form: PageServerData['form']; action: string } = $props()
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
	action={props.action}
	use:enhance
	class="flex gap-4 items-center"
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
	<Form.Button>Create</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
