<script lang="ts">
	import { Drawer as DrawerPrimitive } from 'vaul-svelte'

	let {
		shouldScaleBackground = true,
		open = $bindable(false),
		activeSnapPoint = $bindable(null),
		...restProps
	}: DrawerPrimitive.RootProps = $props()

	function fixBodyPoinerEventsNone() {
		// Using the dropdown menu and the drawer together causes a race condition where the `body`
		// still has `pointer-events: none`, which prevents user interaction with the page. This is a
		// workaround.
		setTimeout(() => {
			// If neither the dropdown nor drawer are open, set it back to `pointer-events: auto`.
			if (
				!document.querySelector('[data-bits-floating-content-wrapper]') &&
				!document.querySelector('[data-vaul-drawer]')
			) {
				document.body.style.pointerEvents = 'auto'
			}
		}, 0)
	}
</script>

<DrawerPrimitive.Root
	{shouldScaleBackground}
	bind:open
	bind:activeSnapPoint
	onAnimationEnd={(open) => {
		if (!open) {
			fixBodyPoinerEventsNone()
		}
	}}
	{...restProps}
/>
