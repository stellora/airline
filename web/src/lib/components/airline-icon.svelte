<script lang="ts">
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		airline,
		class: className,
	}: { airline: Pick<Airline, 'iataCode'>; class?: HTMLAttributes<never>['class'] } = $props()

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

	const colors = airlineColors[airline.iataCode] ?? ['#777', '#ccc']
	const angle =
		180 + (airline.iataCode.split('').reduce((acc, char) => acc + char.charCodeAt(0), 0) % 4) * 45
</script>

<div
	class={cn('w-[1em] h-[1em] rounded-[2px] overflow-hidden', className)}
	role="presentation"
	style="--color-1: {colors[0]}; --color-2: {colors[1]}; --angle: {angle}deg"
>
	<div class="gradient"></div>
</div>

<style>
	.gradient {
		width: 100%;
		height: 100%;
		background: linear-gradient(var(--angle), var(--color-1) 0% 55%, var(--color-2) 80% 100%);
	}
</style>
