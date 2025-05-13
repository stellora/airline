<script lang="ts">
	import type { ComponentProps } from 'svelte'
	import FlightSparkRoute from './flight-spark-route.svelte'
	import ScheduleTitle from './schedule-title.svelte'
	import MileageReward from './mileage-reward.svelte'

	const {
		schedule,
	}: {
		schedule: ComponentProps<typeof ScheduleTitle>['schedule'] &
			ComponentProps<typeof FlightSparkRoute>['flight']
	} = $props()
</script>

<li class="flex flex-col gap-2 border rounded-md min-h-32">
	<div class="flex justify-between items-center pt-3 px-3">
		<ScheduleTitle {schedule} class="text-lg font-bold leading-none" />
		{#if schedule.mileageReward > 0}
			<MileageReward miles={schedule.mileageReward} />
		{/if}
	</div>
	<FlightSparkRoute flight={schedule} />
</li>
