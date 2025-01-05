<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { Itinerary } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		itinerary,
		link = false,
		class: className,
	}: {
		itinerary: Pick<Itinerary, 'recordID'>
		link?: boolean
		tooltip?: boolean
		class?: HTMLAttributes<never>['class']
	} = $props()
</script>

{#if link}
	<a
		href={route('/admin/itineraries/[itinerarySpec]', {
			params: { itinerarySpec: itinerary.recordID },
		})}
		class={cn('font-mono', className)}
	>
		{itinerary.recordID}
	</a>
{:else}
	<span class={cn('font-mono', className)}>{itinerary.recordID}</span>
{/if}
