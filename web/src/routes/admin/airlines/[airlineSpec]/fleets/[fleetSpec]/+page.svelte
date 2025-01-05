<script lang="ts">
	import * as Card from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page title={`Fleet ${data.fleet.code}`}>
	<div class="flex flex-wrap-reverse gap-4">
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>
					{data.fleet.code}
				</Card.Title>
				<Card.Description>
					{data.fleet.description}
				</Card.Description>
			</Card.Header>
			<Card.Content />
		</Card.Root>
	</div>
</Page>
