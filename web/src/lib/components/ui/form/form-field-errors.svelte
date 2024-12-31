<script lang="ts">
	import { cn } from '$lib/utils.js'
	import type { WithoutChild } from 'bits-ui'
	import * as FormPrimitive from 'formsnap'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'
	import { Alert } from '../alert'
	import AlertDescription from '../alert/alert-description.svelte'

	let {
		ref = $bindable(null),
		class: className,
		errorClasses,
		children: childrenProp,
		...restProps
	}: WithoutChild<FormPrimitive.FieldErrorsProps> & {
		errorClasses?: string | undefined | null
	} = $props()
</script>

<FormPrimitive.FieldErrors
	bind:ref
	class={cn('text-destructive text-sm font-medium', className)}
	{...restProps}
>
	{#snippet children({ errors, errorProps })}
		{#if childrenProp}
			{@render childrenProp({ errors, errorProps })}
		{:else}
			{#each errors as error}
				<Alert {...errorProps} variant="destructive" size="sm" class={cn(errorClasses, 'w-fit')}>
					<CircleAlert class="size-5" />
					<AlertDescription>{error}</AlertDescription>
				</Alert>
			{/each}
		{/if}
	{/snippet}
</FormPrimitive.FieldErrors>
