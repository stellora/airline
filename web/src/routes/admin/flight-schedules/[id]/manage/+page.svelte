<script lang="ts">
	import { enhance } from '$app/forms'
	import FlightTitle from '$lib/components/flight-title.svelte'
	import { Button } from '$lib/components/ui/button'
	import { Card, CardContent } from '$lib/components/ui/card'
	import Page from '$lib/components/ui/page/page.svelte'
	import { flightTitle } from '$lib/flight-helpers'
	import NewFlightForm from '../../new/new-flight-form.svelte'

	let { data } = $props()
</script>

<Page title={`Manage ${flightTitle(data.flight)}`}>
	{#snippet titleElement(className)}
		<FlightTitle
			flight={data.flight}
			prefix="Manage "
			class={className}
			subtitleClass="text-base"
			as="h1"
		/>
	{/snippet}
	{#snippet titleActions()}
		<form method="POST" action="?/setFlightSchedulePublished" use:enhance>
			<input type="hidden" name="id" value={data.flight.id} />
			<input type="hidden" name="published" value={data.flight.published ? 'false' : 'true'} />
			<Button type="submit" variant={data.flight.published ? 'outline' : 'default'}>
				{data.flight.published ? 'Unpublish' : 'Publish'}
			</Button>
		</form>
		<form
			method="POST"
			action="?/delete"
			use:enhance={({ cancel }) => {
				if (!confirm('Really delete?')) {
					cancel()
				}
			}}
		>
			<input type="hidden" name="id" value={data.flight.id} />
			<Button type="submit" variant="destructive">Delete flight...</Button>
		</form>
	{/snippet}

	<Card class="self-start">
		<CardContent>
			<NewFlightForm form={data.form} action="?/update" />
		</CardContent>
	</Card>
</Page>
