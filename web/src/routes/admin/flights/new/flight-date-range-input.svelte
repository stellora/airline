<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button'
	import * as Popover from '$lib/components/ui/popover/index.js'
	import { RangeCalendar } from '$lib/components/ui/range-calendar/index.js'
	import { cn } from '$lib/utils.js'
	import { DateFormatter, getLocalTimeZone, today, type DateValue } from '@internationalized/date'
	import CalendarIcon from 'lucide-svelte/icons/calendar'

	const df = new DateFormatter('en-US', {
		dateStyle: 'medium',
	})

	const start = today(getLocalTimeZone())
	const end = start.add({ months: 4 })

	let { value = $bindable() }: { value: { start: DateValue; end: DateValue } } = $props()
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
			{df.format(value.start.toDate(getLocalTimeZone()))} – {df.format(
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
