<script module lang="ts">
	import type { ComponentType } from 'svelte'

	export type PageNavbarTab = {
		title: string
		url: string
		icon: ComponentType
	}
</script>

<script lang="ts">
	import { page } from '$app/state'
	import Button from '../button/button.svelte'

	const {
		tabs,
	}: {
		tabs: PageNavbarTab[]
	} = $props()
</script>

<nav class="flex gap-2">
	{#each tabs as tab (tab.url)}
		<Button
			variant="pageNavbarTab"
			size="pageNavbar"
			href={tab.url}
			data-active={page.url.pathname === tab.url}
		>
			<tab.icon />
			{tab.title}</Button
		>
	{/each}
</nav>
