<script lang="ts">
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import FlightList from '../../../../../admin/flights/flight-list.svelte'

	let { data } = $props()
</script>

<Page title={`Flights for scheduled ${flightTitle(data.schedule)}`}>
	<ScheduleTitle flight={data.schedule} as="span" showRoute={true} />
	<FlightList flights={data.flights} />
</Page>
