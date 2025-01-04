<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements'
	import Input from './ui/input/input.svelte'

	let { value = $bindable(undefined), ...restProps }: Omit<HTMLInputAttributes, 'type'> = $props()
</script>

<div class="time-of-day-input">
	<Input bind:value type="time" {...restProps} />
</div>

<style>
	.time-of-day-input > :global(input) {
		/* Needed to avoid AM/PM being clipped if the user input starts with "00:". */
		@apply font-mono;
	}
	.time-of-day-input > :global(input::-webkit-datetime-edit-ampm-field) {
		opacity: 60%;

		margin-left: -0.25em;
	}
</style>
