<script lang="ts">
	import { page } from '$app/state'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { route } from '$lib/route-helpers'
	import FlightInstanceList from '../../../admin/flight-instances/flight-instance-list.svelte'

	let { data } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]/flights', {
				params: { airlineSpec: page.params.airlineSpec },
			})}>Flights</Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

<Page title={`Flights - ${data.airline.iataCode}`}>
	<FlightInstanceList flightInstances={data.flightInstances!} showFlightInfo />
</Page>
