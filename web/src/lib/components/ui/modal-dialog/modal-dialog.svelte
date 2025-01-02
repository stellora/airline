<script module>
	import { pushState, replaceState } from '$app/navigation'

	export function openModalDialog(id: string): void {
		pushState('', { showModal: id })
	}
</script>

<script lang="ts">
	import { page } from '$app/state'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import type { Snippet } from 'svelte'

	const { id, children }: { id: string; children: Snippet<[]> } = $props()
</script>

<Dialog.Root
	bind:open={() => page.state.showModal === id,
	(open) => {
		if (open) {
			pushState('', { showModal: id })
		} else {
			replaceState('', {})
		}
	}}
>
	<Dialog.Trigger>Open</Dialog.Trigger>
	<Dialog.Content>{@render children()}</Dialog.Content>
</Dialog.Root>
