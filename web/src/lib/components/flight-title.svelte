<script lang="ts">
	import type { Flight } from '$lib/types'
	import { cn } from '$lib/utils'
	import Star from 'lucide-svelte/icons/star'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		flight,
		link = false,
		class: className,
		as = 'h2'
	}: {
		flight: Flight
		link?: boolean
		class: HTMLAttributes<never>['class']
		as?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'span'
	} = $props()
</script>

<svelte:element this={as} class={cn(className, 'flex items-center gap-1.5')}>
	{#if link}
		<a href={`/admin/flights/${flight.id}`}>{flight.title}</a>
	{:else}
		{flight.title}
	{/if}
	{#if flight.published}
		<Star fill="currentColor" class="inline" size="1.25ch" aria-label="Published" />
	{/if}
</svelte:element>
