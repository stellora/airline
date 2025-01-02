<script lang="ts">
	import type { Aircraft } from '$lib/types'

	const { aircraft }: { aircraft: Pick<Aircraft, 'registration'> } = $props()
</script>

<span class="font-mono">{aircraft.registration}</span>
