<script lang="ts">
	import type { Aircraft } from '$lib/types'
	import AircraftTypeCode from './aircraft-type-code.svelte'

	const {
		aircraft,
		showAircraftType = false,
	}: { aircraft: Pick<Aircraft, 'registration' | 'aircraftType'>; showAircraftType?: boolean } =
		$props()
</script>

<div class="inline-flex gap-1">
	<span class="font-mono">{aircraft.registration}</span>
	{#if showAircraftType}
		<AircraftTypeCode aircraftType={aircraft.aircraftType} class="text-muted-foreground" />
	{/if}
</div>
