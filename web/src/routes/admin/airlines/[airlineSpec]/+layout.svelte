<script lang="ts">
	import AirlineCode from '$lib/components/airline-code.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import { route } from '$lib/route-helpers'

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

{@render children()}
