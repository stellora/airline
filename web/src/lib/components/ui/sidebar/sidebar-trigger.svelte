<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js'
	import * as Tooltip from '$lib/components/ui/tooltip'
	import { cn } from '$lib/utils.js'
	import PanelLeftClose from 'lucide-svelte/icons/panel-left-close'
	import PanelLeftOpen from 'lucide-svelte/icons/panel-left-open'
	import type { ComponentProps } from 'svelte'
	import { useSidebar } from './context.svelte.js'

	let {
		ref = $bindable(null),
		class: className,
		onclick,
		location,
		...restProps
	}: ComponentProps<typeof Button> & {
		onclick?: (e: MouseEvent) => void
		location: 'navbar'
	} = $props()

	const sidebar = useSidebar()
</script>

{#if !(sidebar.isMobile ? sidebar.openMobile : sidebar.open)}
	<Tooltip.Root>
		<Tooltip.Trigger>
			{#snippet child({ props })}
				<Button
					{...props}
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
					{#if location === 'navbar'}
						<PanelLeftOpen />
					{:else if location === 'sidebar'}
						<PanelLeftClose />
					{/if}
					<span class="sr-only">Toggle Sidebar</span>
				</Button>
			{/snippet}
		</Tooltip.Trigger>
		<Tooltip.Content>{sidebar.open ? 'Close' : 'Open'} sidebar</Tooltip.Content>
	</Tooltip.Root>
{/if}
