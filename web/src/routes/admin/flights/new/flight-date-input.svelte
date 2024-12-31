<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button/index.js'
	import { Calendar } from '$lib/components/ui/calendar/index.js'
	import * as Popover from '$lib/components/ui/popover/index.js'
	import { cn } from '$lib/utils.js'
	import { DateFormatter, type DateValue, getLocalTimeZone } from '@internationalized/date'
	import CalendarIcon from 'lucide-svelte/icons/calendar'

	// TODO!(sqs): not used

	const df = new DateFormatter('en-US', {
		dateStyle: 'long',
	})

	let { value = $bindable(undefined) }: { value?: DateValue | undefined } = $props()
	let contentRef = $state<HTMLElement | null>(null)
</script>

<Popover.Root>
	<Popover.Trigger
		class={cn(
			'!flex',
			buttonVariants({
				variant: 'outline',
				class: 'w-[210px] justify-start text-left font-normal',
			}),
			!value && 'text-muted-foreground',
		)}
	>
		<CalendarIcon />
		{value ? df.format(value.toDate(getLocalTimeZone())) : 'Date...'}
	</Popover.Trigger>
	<Popover.Content bind:ref={contentRef} class="w-auto p-0" align="start">
		<Calendar type="single" bind:value />
	</Popover.Content>
</Popover.Root>
