<script lang="ts">
	import { globalAdminItems } from '$lib/components/app-sidebar.svelte'
	import Button from '$lib/components/ui/button/button.svelte'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title="Global admin">
	<nav class="grid grid-cols-[repeat(auto-fill,minmax(225px,1fr))] gap-6">
		{#each globalAdminItems as item (item.title)}
			<Button variant="outline" href={item.url} class="h-[200px] text-xl"
				><item.icon class="!size-7 [&:not(:hover)]:text-muted-foreground" /> {item.title}</Button
			>
		{/each}
	</nav>
</Page>
