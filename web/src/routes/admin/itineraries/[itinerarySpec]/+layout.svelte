<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import Trash from 'lucide-svelte/icons/trash'

	const { data, children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/admin/itineraries/[itinerarySpec]', {
				params: { itinerarySpec: data.itinerary.recordID },
			})}>{data.itinerary.recordID}</Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav>
	{#snippet breadcrumbActions()}
		<PageNavbarBreadcrumbActionsDropdownMenu>
			<DropdownMenu.Group>
				<DropdownMenu.Item>
					{#snippet child({ props })}
						<form
							method="POST"
							action={route('/admin/itineraries/[itinerarySpec]', {
								params: { itinerarySpec: page.params.itinerarySpec },
								query: '/delete',
							})}
							use:enhance={({ cancel }) => {
								if (!confirm('Really delete?')) {
									cancel()
								}
							}}
							class="w-full [&>button]:w-full"
						>
							<button type="submit" {...props}>
								<Trash /> Delete...
							</button>
						</form>
					{/snippet}</DropdownMenu.Item
				>
			</DropdownMenu.Group>
		</PageNavbarBreadcrumbActionsDropdownMenu>
	{/snippet}
</PageNav>

{@render children()}
