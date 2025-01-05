<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	const { children } = $props()
</script>

<Portal target="[aria-label=breadcrumb]">HELLO123</Portal>

{@render children?.()}
