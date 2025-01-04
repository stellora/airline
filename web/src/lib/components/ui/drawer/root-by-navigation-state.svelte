<script lang="ts">
	import { pushState } from '$app/navigation'
	import { page } from '$app/state'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import type { ComponentProps } from 'svelte'

	const { id, ...restProps }: { id: string } & ComponentProps<typeof Drawer.Root> = $props()
</script>

<Drawer.Root
	handleOnly
	modal
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
