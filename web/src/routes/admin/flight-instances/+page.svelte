<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import FlightInstanceList from './flight-instance-list.svelte'

	let { data } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link href={route('/admin/flight-instances')}>Flights</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<Page title="Flight instances">
	<FlightInstanceList flightInstances={data.flightInstances} showFlightInfo />
</Page>
