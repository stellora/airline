<script lang="ts">
	import { route } from '$lib/route-helpers'
	import type { Aircraft } from '$lib/types'
	import AircraftTypeCode from './aircraft-type-code.svelte'

	const {
		aircraft,
		link = false,
		showAircraftType = false,
	}: {
		aircraft: Pick<Aircraft, 'registration' | 'aircraftType'>
		link?: boolean
		showAircraftType?: boolean
	} = $props()
</script>

<div class="inline-flex gap-1">
	{#if link}
		<a href={route('/admin/aircraft/[aircraftSpec]')} class="font-mono">{aircraft.registration}</a>
	{:else}
		<span class="font-mono">{aircraft.registration}</span>
	{/if}
	{#if showAircraftType}
		<AircraftTypeCode aircraftType={aircraft.aircraftType} class="text-muted-foreground" />
	{/if}
</div>
