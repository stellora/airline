<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import Ellipsis from 'lucide-svelte/icons/ellipsis'
	import type { Snippet } from 'svelte'
	import { buttonVariants } from '../button'

	const {
		children,
	}: {
		children: Snippet<[]>
	} = $props()
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger
		class={buttonVariants({
			variant: 'pageNavbarIcon',
			size: 'pageNavbar',
		})}><Ellipsis /></DropdownMenu.Trigger
	>
	<DropdownMenu.Content>
		{@render children()}
	</DropdownMenu.Content>
</DropdownMenu.Root>
