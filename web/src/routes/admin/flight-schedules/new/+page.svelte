<script lang="ts">
	import { Card, CardContent } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import FlightScheduleForm from '../flight-schedule-form.svelte'

	let { data } = $props()

	// TODO!(sqs): for testing +page.svelte files directly: https://github.com/sveltejs/kit/issues/5525#issuecomment-1186390654
</script>

<Page title="New flight schedule">
	<Card>
		<CardContent>
			<FlightScheduleForm form={data.form} action="?/create" />
		</CardContent>
	</Card>
</Page>
