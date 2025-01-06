<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { Fleet, Itinerary } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		fleet,
		link = false,
		class: className,
	}: {
		fleet: Pick<Fleet, 'code'> & { airline: Pick<Fleet['airline'], 'iataCode'> }
		link?: boolean
		tooltip?: boolean
		class?: HTMLAttributes<never>['class']
	} = $props()
</script>

{#if link}
	<a
		href={route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
			params: { airlineSpec: fleet.airline.iataCode, fleetSpec: fleet.code },
		})}
		class={cn('font-mono', className)}
	>
		{fleet.code}
	</a>
{:else}
	<span class={cn('font-mono', className)}>{fleet.code}</span>
{/if}
