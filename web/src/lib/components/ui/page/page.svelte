<script lang="ts">
	import type { Snippet } from 'svelte'
	import type { DOMAttributes } from 'svelte/elements'

	const {
		title,
		titleElement,
		children,
	}: { title: string; titleElement?: Snippet<[className: string]> } & Pick<
		DOMAttributes<any>,
		'children'
	> = $props()
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div class="flex flex-col gap-4">
	{#if titleElement}
		{@render titleElement('text-4xl font-bold')}
	{:else}
		<h1 class="text-4xl font-bold">{title}</h1>
	{/if}

	{@render children?.()}
</div>
