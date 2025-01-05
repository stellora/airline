<script lang="ts">
	import { page } from '$app/state'
	import { schema } from '$lib/airline.typebox'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Drawer from '$lib/components/ui/drawer/index.js'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import { route } from '$lib/route-helpers'
	import Settings_2 from 'lucide-svelte/icons/settings-2'
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
	{#snippet actions()}
		<Drawer.DrawerByNavigationState id="edit-airport" direction="right">
			<Drawer.Trigger class={buttonVariants({ variant: 'secondary', size: 'pageNavbar' })}>
				<Settings_2 /> Edit
			</Drawer.Trigger>
			<Drawer.Content>
				<Drawer.Header>
					<Drawer.Title>Edit airport</Drawer.Title>
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
</PageNav>

{@render children()}
