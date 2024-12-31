<script lang="ts">
	import { Card, CardContent } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import NewFlightForm from './new-flight-form.svelte'

	let { data } = $props()

	// TODO!(sqs): for testing +page.svelte files directly: https://github.com/sveltejs/kit/issues/5525#issuecomment-1186390654
</script>

<Page title="New flight">
	<Card class="self-start">
		<CardContent>
			<NewFlightForm form={data.form} />
		</CardContent>
	</Card>
</Page>
