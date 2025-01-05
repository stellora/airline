<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements'
	import Input from './ui/input/input.svelte'

	let {
		value = $bindable(undefined),
		...restProps
	}: { value: number | undefined } & Omit<HTMLInputAttributes, 'type' | 'value'> = $props()

	let tmpInput = $state<string | undefined>()
</script>

<Input
	type="text"
	pattern="\d?\d:\d\d"
	bind:value={() => {
		if (tmpInput !== undefined) {
			return tmpInput
		}
		const hours = Math.floor((value ?? 0) / 3600)
		const minutes = Math.floor(((value ?? 0) % 3600) / 60)
		return `${hours}:${minutes.toString().padStart(2, '0')}`
	},
	(v) => {
		const match = v.match(/^(\d?\d):(\d\d)$/)
		if (match) {
			const hours = parseInt(match[1])
			const minutes = parseInt(match[2])
			value = hours * 3600 + minutes * 60
			tmpInput = undefined
		} else {
			value = undefined
			tmpInput = v
		}
	}}
	{...restProps}
/>
