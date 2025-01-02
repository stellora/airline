<script lang="ts">
	import type { Airline } from '$lib/types'

	const { airline }: { airline: Pick<Airline, 'iataCode'> } = $props()

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
	class="w-[1rem] h-[1rem] relative rounded-[2px] overflow-hidden"
	role="presentation"
	style="--color-1: {colors[0]}; --color-2: {colors[1]}; --angle: {angle}deg"
>
	<div class="gradient"></div>
</div>

<style>
	.gradient {
		position: absolute;
		inset: 0;
		background: linear-gradient(var(--angle), var(--color-1) 0% 55%, var(--color-2) 80% 100%);
	}
</style>
