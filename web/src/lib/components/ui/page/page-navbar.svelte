<script module>
	export const PAGE_NAVBAR_PORTAL_ID = '#page-navbar-portal'
	export const PAGE_BREADCRUMBS_PORTAL_ID = '#page-breadcrumbs-portal'
</script>

<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import * as Sidebar from '$lib/components/ui/sidebar'
	import { buttonVariants } from '../button'
</script>

<div
	class="flex items-center overflow-x-auto overflow-y-hidden gap-2 px-4 py-2 sticky top-0 border-b bg-background h-[47px]"
>
	<Sidebar.Trigger
		location="navbar"
		class={buttonVariants({
			variant: 'pageNavbarIcon',
			size: 'pageNavbar',
			class: '[&>svg]:!size-4 px-2',
		})}
	/>
	<Breadcrumb.Root
		class="[&>ol>li:not(:last-child)]:max-md:hidden [&>ol>li:last-child[role=presentation]]:hidden"
	>
		<Breadcrumb.List id={PAGE_BREADCRUMBS_PORTAL_ID.slice(1)} />
	</Breadcrumb.Root>
	<div id={PAGE_NAVBAR_PORTAL_ID.slice(1)} class="flex-1">
		<!-- Use <Portal target={PAGE_NAVBAR_BREADCRUMB_ACTIONS_ID}> to supply content. -->
	</div>
</div>
