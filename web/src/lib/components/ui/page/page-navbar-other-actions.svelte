<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import Ellipsis from 'lucide-svelte/icons/ellipsis'
	import { buttonVariants } from '../button'

	const {
		tabs,
	}: {
		tabs: PageNavbarTabsItem[]
	} = $props()
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger
		class={buttonVariants({
			variant: 'ghost',
			size: 'sm',
			class:
				'h-[unset] py-1 px-1 text-muted-foreground hover:text-sidebar-accent-foreground [&>svg]:!size-5',
		})}><Ellipsis /></DropdownMenu.Trigger
	>
	<DropdownMenu.Content>
		<DropdownMenu.Group>
			<DropdownMenu.GroupHeading>My Account</DropdownMenu.GroupHeading>
			<DropdownMenu.Separator />
			<DropdownMenu.Item>Profile</DropdownMenu.Item>
			<DropdownMenu.Item>Billing</DropdownMenu.Item>
			<DropdownMenu.Item>Team</DropdownMenu.Item>
			<DropdownMenu.Item>Subscription</DropdownMenu.Item>
		</DropdownMenu.Group>
	</DropdownMenu.Content>
</DropdownMenu.Root>
