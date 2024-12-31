<script lang="ts">
	import Page from '$lib/components/ui/page/page.svelte'
	import FlightForm from './flight-form.svelte'

	let { data } = $props()

	// TODO!(sqs): for testing +page.svelte files directly: https://github.com/sveltejs/kit/issues/5525#issuecomment-1186390654
</script>

<Page title="New flight">
	<FlightForm form={data.form} />
</Page>
