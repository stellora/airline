<script lang="ts">
	import type { Flight } from '$lib/types'
	import type { HTMLAttributes } from 'svelte/elements'

	const {
		flight,
		class: className,
	}: {
		flight: Flight
		class?: HTMLAttributes<never>['class']
	} = $props()
</script>

<span>Scheduled</span>
