<script lang="ts">
	import FlightTitle from '$lib/components/flight-title.svelte'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { route } from '$lib/route-helpers'
	import type { FlightSchedule } from '$lib/types'

	let { flightSchedule }: { flightSchedule: FlightSchedule } = $props()
</script>

<Breadcrumb.Item
	><Breadcrumb.Link
		href={route('/manage/[airlineSpec]/schedules/[id]', {
			params: { airlineSpec: flightSchedule.airline.iataCode, id: flightSchedule.id.toString() },
		})}
		class="[&_[data-airline-icon]]:-mt-[1.7px] mt-[2.2px] block"
		><FlightTitle
			flight={flightSchedule}
			as="span"
			showRoute={true}
			link={true}
			tooltip={false}
		/></Breadcrumb.Link
	></Breadcrumb.Item
>
