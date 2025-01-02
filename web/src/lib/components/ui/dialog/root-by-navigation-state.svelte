<script lang="ts">
	import { pushState } from '$app/navigation'
	import { page } from '$app/state'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import type { ComponentProps } from 'svelte'

	const { id, ...restProps }: { id: string } & ComponentProps<typeof Dialog.Root> = $props()
</script>

<Dialog.Root
	bind:open={() => page.state.showModal === id,
	(open) => {
		if (open) {
			pushState('', { showModal: id })
		} else {
			history.back()
		}
	}}
	{...restProps}
/>
