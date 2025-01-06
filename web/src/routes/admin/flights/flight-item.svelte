<script lang="ts">
	import ScheduleTitle from '$lib/components/schedule-title.svelte'
	import type { ComponentProps } from 'svelte'
	import { fade } from 'svelte/transition'

	const {
		flight,
	}: {
		flight: ComponentProps<typeof ScheduleTitle>['schedule']
	} = $props()
</script>

<li in:fade out:fade class="flex flex-col gap-4 border p-3 rounded-md">
	<ScheduleTitle schedule={flight} link class="text-lg font-bold leading-none" />
</li>
