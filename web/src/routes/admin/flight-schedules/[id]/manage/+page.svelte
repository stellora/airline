<script lang="ts">
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Card, CardContent } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import FlightScheduleForm from '../../flight-schedule-form.svelte'

	let { data } = $props()
</script>

<Page title={`Manage ${flightTitle(data.flightSchedule)}`}>
	{#snippet titleElement(className)}
		<FlightTitle
			flight={data.flightSchedule}
			prefix="Manage "
			class={className}
			subtitleClass="text-base"
			as="h1"
		/>
	{/snippet}
	<Card class="self-start">
		<CardContent>
			<FlightScheduleForm form={data.form} action="?/update" />
		</CardContent>
	</Card>
</Page>
