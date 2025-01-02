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

<nav class="flex flex-wrap gap-2">
	{#each tabs as tab (tab.url)}
		<Button
			variant="outline"
			size="sm"
			class="h-[unset] py-1 px-2 text-muted-foreground [&:not(:hover)>svg]:text-muted-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground active:bg-sidebar-accent active:text-sidebar-accent-foreground data-[active=true]:bg-sidebar-accent data-[active=true]:text-sidebar-accent-foreground [&>svg]:size-4"
			href={tab.url}
			data-active={page.url.pathname === tab.url}
		>
			<tab.icon />
			{tab.title}</Button
		>
	{/each}
</nav>
