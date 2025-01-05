<script lang="ts">
	import type { FlightInstance } from '$lib/types'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		flight,
		class: className,
	}: {
		flight: FlightInstance
		class?: HTMLAttributes<never>['class']
	} = $props()
</script>

<span>Scheduled</span>
