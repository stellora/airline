<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import { formatDateFull } from '$lib/datetime-helpers'
	import { cn } from '$lib/utils'
	import type { ZonedDateTime } from '@internationalized/date'
	import type { Snippet } from 'svelte'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		value,
		children,
		class: className,
	}: {
		value: ZonedDateTime
		children: Snippet
		class?: HTMLAttributes<never>['class']
	} = $props()
</script>

<Tooltip.Root>
	<Tooltip.Trigger>
		{#snippet child({ props })}
			<time {...props} class={cn('cursor-help', className)} datetime={value.toString()}>
				{@render children()}
			</time>
		{/snippet}
	</Tooltip.Trigger><Tooltip.Portal>
		<Tooltip.Content>
			{formatDateFull(value)}
		</Tooltip.Content>
	</Tooltip.Portal>
</Tooltip.Root>
