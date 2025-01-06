<script lang="ts">
	import { enhance } from '$app/forms'
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import PageNavbarBreadcrumbActionsDropdownMenu from '$lib/components/ui/page/page-navbar-breadcrumb-actions-dropdown-menu.svelte'
	import { route } from '$lib/route-helpers'
	import { cn } from '$lib/utils'
	import Settings_2 from 'lucide-svelte/icons/settings-2'
	import Trash from 'lucide-svelte/icons/trash'
	import type { ClassNameValue } from 'tailwind-merge'
	import AirlineForm from '../airline-form.svelte'

	const { data, children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/admin/airlines/[airlineSpec]', {
				params: { airlineSpec: data.airline.iataCode },
			})}><AirlineCode airline={data.airline} icon /></Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<PageNav>
	{#snippet breadcrumbActions()}
		<Drawer.DrawerByNavigationState id="edit-airline" direction="right">
			<PageNavbarBreadcrumbActionsDropdownMenu>
				<DropdownMenu.Group>
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<Drawer.Trigger {...props} class={cn(props.class as ClassNameValue, 'w-full')}>
								<Settings_2 /> Edit
							</Drawer.Trigger>
						{/snippet}
					</DropdownMenu.Item>
					<DropdownMenu.Separator />
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<form
								method="POST"
								action={route('/admin/airlines/[airlineSpec]', {
									params: { airlineSpec: page.params.airlineSpec },
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
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit airline</Drawer.Title>
				</Drawer.Header>
				<Drawer.ScrollArea>
					<AirlineForm
						action={route('/admin/airlines/[airlineSpec]', {
							params: { airlineSpec: page.params.airlineSpec },
							query: '/update',
						})}
						submitLabel="Save"
						data={data.form}
						schema={schema['/airlines/{airlineSpec}']['PATCH']['args']['properties']['body']}
					/>
				</Drawer.ScrollArea>
			</Drawer.Content>
		</Drawer.DrawerByNavigationState>
	{/snippet}
	{#snippet actions()}
		<Button
			variant="outline"
			size="pageNavbar"
			href={route('/manage/[airlineSpec]', {
				params: { airlineSpec: data.airline.iataCode },
			})}>Switch to management view</Button
		>
	{/snippet}
</PageNav>

{@render children()}
