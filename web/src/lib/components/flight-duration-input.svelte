<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements'
	import Input from './ui/input/input.svelte'

	let {
		value = $bindable(undefined),
		...restProps
	}: { value: number | undefined } & Omit<HTMLInputAttributes, 'type' | 'value'> = $props()
</script>

<div class="inline-flex gap-2">
	<Input bind:value type="text" pattern="\d?\d:\d\d" {...restProps} />
</div>
=
