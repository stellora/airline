<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Popover from '$lib/components/ui/popover/index.js'
	import { RangeCalendar } from '$lib/components/ui/range-calendar/index.js'
	import { dateFormatter } from '$lib/datetime-helpers'
	import { cn } from '$lib/utils.js'
	import { getLocalTimeZone } from '@internationalized/date'
	import type { DateRange } from 'bits-ui'
	import CalendarIcon from 'lucide-svelte/icons/calendar'

	let { value = $bindable() }: { value: DateRange } = $props()
	let contentRef = $state<HTMLElement | null>(null)
</script>

<Popover.Root>
	<Popover.Trigger
		class={cn(
			'!flex',
			buttonVariants({
				variant: 'outline',
				class: 'justify-start text-left font-normal',
			}),
			!value && 'text-muted-foreground',
		)}
	>
		<CalendarIcon />
		{#if value && value.start && value.end}
			{dateFormatter.format(value.start.toDate(getLocalTimeZone()))} – {dateFormatter.format(
				value.end.toDate(getLocalTimeZone()),
			)}
		{:else}
			Select date range...
		{/if}
	</Popover.Trigger>
	<Popover.Content bind:ref={contentRef} class="w-auto p-0" align="start">
		<RangeCalendar bind:value class="rounded-md border" />
	</Popover.Content>
</Popover.Root>
