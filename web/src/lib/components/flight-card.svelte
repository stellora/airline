<script lang="ts">
	import type { ComponentProps } from 'svelte'
	import FlightSparkRoute from './flight-spark-route.svelte'
	import ScheduleTitle from './schedule-title.svelte'

	const {
		schedule,
	}: {
		schedule: ComponentProps<typeof ScheduleTitle>['schedule'] &
			ComponentProps<typeof FlightSparkRoute>['flight']
	} = $props()
</script>

<li class="flex flex-col gap-2 border rounded-md min-h-32">
	<ScheduleTitle {schedule} class="text-lg font-bold leading-none pt-3 px-3" />
	<FlightSparkRoute flight={schedule} />
</li>
