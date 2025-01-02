<script lang="ts">
	import Portal from '$lib/components/portal.svelte'
	import type { Snippet } from 'svelte'
	import { PAGE_NAVBAR_ACTIONS_ID } from './page-navbar.svelte'

	const {
		breadcrumbActions,
		title,
		showTitleHeading = false,
		titleElement,
		titleActions,
		children,
	}: {
		breadcrumbActions?: Snippet
		title: string
		showTitleHeading?: boolean
		titleElement?: Snippet<[className: string]>
		titleActions?: Snippet // TODO!(sqs): remove
		children?: Snippet
	} = $props()
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

{#if breadcrumbActions}
	<Portal target={PAGE_NAVBAR_ACTIONS_ID}>
		{@render breadcrumbActions()}
	</Portal>
{/if}

<main class="flex flex-col gap-4">
	{#if showTitleHeading}
		<header class="flex flex-wrap gap-4 justify-between items-start">
			{#if titleElement}
				{@render titleElement('text-4xl font-bold')}
			{:else}
				<h1 class="text-4xl font-bold">{title}</h1>
			{/if}
			<div class="flex-1"></div>
			{@render titleActions?.()}
		</header>
	{/if}

	{@render children?.()}
</main>
