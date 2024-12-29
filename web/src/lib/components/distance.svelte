<script lang="ts">
	let { distanceMiles }: { distanceMiles?: number | undefined } = $props()
	if (distanceMiles !== undefined) {
		distanceMiles = Math.round(distanceMiles)
	}
</script>

{#if distanceMiles !== undefined}
	{Math.round(distanceMiles)}{' '}{distanceMiles === 1 ? 'mile' : 'miles'}
{/if}
