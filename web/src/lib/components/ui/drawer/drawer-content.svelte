<script lang="ts">
	import { cn } from '$lib/utils.js'
	import X from 'lucide-svelte/icons/x'
	import { Drawer as DrawerPrimitive } from 'vaul-svelte'
	import { DrawerClose } from '.'
	import DrawerOverlay from './drawer-overlay.svelte'

	let {
		ref = $bindable(null),
		class: className,
		portalProps,
		children,
		...restProps
	}: DrawerPrimitive.ContentProps & {
		portalProps?: DrawerPrimitive.PortalProps
	} = $props()
</script>

<DrawerPrimitive.Portal {...portalProps}>
	<DrawerOverlay />
	<DrawerPrimitive.Content
		bind:ref
		class={cn(
			'bg-background fixed inset-y-0 right-0 z-50 w-[80%] max-w-[650px] flex flex-col border max-h-screen overflow-hidden',
			className,
		)}
		{...restProps}
	>
		{@render children?.()}
		<DrawerClose
			class="ring-offset-background focus:ring-ring absolute right-4 top-4 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:pointer-events-none"
		>
			<X class="size-4" />
			<span class="sr-only">Close</span>
		</DrawerClose>
	</DrawerPrimitive.Content>
</DrawerPrimitive.Portal>
