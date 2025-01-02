<script lang="ts">
	import AircraftRegistration from '$lib/components/aircraft-registration.svelte'
	import AircraftTypeCode from '$lib/components/aircraft-type-code.svelte'
	import AirlineCode from '$lib/components/airline-code.svelte'
	import Page from '$lib/components/ui/page/page.svelte'

	let { data } = $props()
</script>

<Page
	title={`${data.aircraft.registration} (${data.aircraft.airline.iataCode} ${data.aircraft.aircraftType})}`}
	showTitleHeading
>
	{#snippet titleElement()}
		<div class="flex items-baseline gap-2">
			<h1 class="text-4xl font-bold">
				<AircraftRegistration aircraft={data.aircraft} />
			</h1>
			<span class="text-muted-foreground flex items-center gap-2">
				<AircraftTypeCode aircraftType={data.aircraft.aircraftType} />
				<AirlineCode airline={data.aircraft.airline} icon={true} link={true} />
			</span>
		</div>
	{/snippet}
</Page>
