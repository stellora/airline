<script lang="ts">
	import { pushState } from '$app/navigation'
	import { page } from '$app/state'
	import * as Sheet from '$lib/components/ui/sheet/index.js'
	import type { ComponentProps } from 'svelte'

	const { id, ...restProps }: { id: string } & ComponentProps<typeof Sheet.Root> = $props()
</script>

<Sheet.Root
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
