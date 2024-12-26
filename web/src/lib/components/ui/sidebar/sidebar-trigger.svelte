<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js'
	import { cn } from '$lib/utils.js'
	import PanelLeft from 'lucide-svelte/icons/panel-left'
	import type { ComponentProps } from 'svelte'
	import { useSidebar } from './context.svelte.js'

	let {
		ref = $bindable(null),
		class: className,
		onclick,
		mobileOnly,
		...restProps
	}: ComponentProps<typeof Button> & {
		onclick?: (e: MouseEvent) => void
		mobileOnly?: boolean
	} = $props()

	const sidebar = useSidebar()
</script>

{#if !mobileOnly || sidebar.isMobile || !sidebar.open}
	<Button
		type="button"
		onclick={(e) => {
			onclick?.(e)
			sidebar.toggle()
		}}
		data-sidebar="trigger"
		variant="ghost"
		size="icon"
		class={cn('h-7 w-7', className)}
		{...restProps}
	>
		<PanelLeft />
		<span class="sr-only">Toggle Sidebar</span>
	</Button>
{/if}
