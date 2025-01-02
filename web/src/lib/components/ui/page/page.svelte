<script lang="ts">
	import type { Snippet } from 'svelte'

	const {
		title,
		showTitleHeading = false,
		titleElement,
		children,
	}: {
		title: string
		showTitleHeading?: boolean
		titleElement?: Snippet<[className: string]>
		children?: Snippet
	} = $props()
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<main class="flex flex-col gap-4">
	{#if showTitleHeading}
		<header class="flex flex-wrap gap-4 justify-between items-start">
			{#if titleElement}
				{@render titleElement('text-4xl font-bold')}
			{:else}
				<h1 class="text-4xl font-bold">{title}</h1>
			{/if}
		</header>
	{/if}

	{@render children?.()}
</main>
