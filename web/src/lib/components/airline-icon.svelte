<script lang="ts">
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		airline,
		size = 'default',
		showCode = false,
		class: className,
	}: {
		airline: Pick<Airline, 'iataCode'>
		size?: 'default' | 'lg'
		showCode?: boolean
		class?: HTMLAttributes<never>['class']
	} = $props()

	const airlineColors: Record<Airline['iataCode'], [color1: string, color2: string]> = {
		AC: ['#D22630', '#000000'],
		AA: ['#0078D2', '#C00C2C'],
		BA: ['#2A5c9a', '#cf1c06'],
		EY: ['#BD8B13', '#D4A12A'],
		UA: ['#0154A6', '#0C2340'],
		SQ: ['#fcb130', '#1d4886'],
		LH: ['#05164D', '#FFAD1D'],
		AF: ['#F71D25', '#002157'],
		JL: ['#FF0000', '#FFFFFF'],
		KL: ['#00A1DE', '#FFFFFF'],
		DL: ['#003268', '#E3132C'],
		LX: ['#E60005', '#FFFFFF'],
		NZ: ['#00247D', '#FFFFFF'],
		QF: ['#EE0000', '#FFFFFF'],
		QR: ['#5C0632', '#FFFFFF'],
		EK: ['#D71A21', '#FFFFFF'],
	}

	const colors = $derived(airlineColors[airline.iataCode] ?? ['#777', '#ccc'])
	const angle = $derived(
		180 + (airline.iataCode.split('').reduce((acc, char) => acc + char.charCodeAt(0), 0) % 4) * 45,
	)
</script>

<div
	class={cn(
		'rounded-[2px] overflow-hidden inline-block relative',
		{
			'w-[1em] h-[1em] rounded-[2px]': size === 'default',
			'w-[2em] h-[2em] rounded-[4px]': size === 'lg',
		},
		className,
	)}
	role="presentation"
	style="--color-1: {colors[0]}; --color-2: {colors[1]}; --angle: {angle}deg"
	data-airline-icon
>
	<div class="gradient"></div>
	{#if showCode}
		<div
			class="absolute inset-0 flex items-center justify-center text-white text-[0.9em] font-bold font-mono"
		>
			{airline.iataCode}
		</div>
	{/if}
</div>

<style>
	.gradient {
		width: 100%;
		height: 100%;
		background: linear-gradient(var(--angle), var(--color-1) 0% 55%, var(--color-2) 80% 100%);
	}
</style>
