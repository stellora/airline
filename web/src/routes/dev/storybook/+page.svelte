<script lang="ts">
	import FlightSparkRoute from '$lib/components/flight-spark-route.svelte'
	import GreatCircleRoute from '$lib/components/maps/great-circle-route.svelte'
	import { FIXTURE_AIRPORTS } from '$lib/fixtures'
</script>

<svelte:head>
	<title>Storybook</title>
</svelte:head>

<h1 class="text-4xl font-bold">Storybook</h1>

<section class="my-6">
	<h2 class="text-2xl font-bold mb-4">GreatCircleRoute</h2>

	<div class="grid grid-cols-[repeat(auto-fill,minmax(350px,1fr))] gap-4 *:border *:rounded">
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.EWR]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.LHR, FIXTURE_AIRPORTS.AMS]]} detailLevel="high" />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.EWR, FIXTURE_AIRPORTS.SIN]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.SCL]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.LHR]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.EWR]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.SIN]]} />
		<GreatCircleRoute
			routes={[
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.LHR],
				[FIXTURE_AIRPORTS.EWR, FIXTURE_AIRPORTS.AMS],
			]}
		/>
		<GreatCircleRoute
			routes={[
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.SIN],
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.EWR],
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.LHR],
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.AMS],
			]}
		/>
		<GreatCircleRoute
			routes={[
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.SIN],
				[FIXTURE_AIRPORTS.LHR, FIXTURE_AIRPORTS.SCL],
			]}
		/>
	</div>
</section>

<section>
	<h2 class="text-2xl font-bold mb-4">FlightSparkRoute</h2>

	<div class="flex flex-wrap gap-4 *:w-[100px] *:h-[25px] *:border *:rounded">
		<FlightSparkRoute
			flight={{ originAirport: FIXTURE_AIRPORTS.SFO, destinationAirport: FIXTURE_AIRPORTS.LHR }}
		/>
		<FlightSparkRoute
			flight={{ originAirport: FIXTURE_AIRPORTS.SFO, destinationAirport: FIXTURE_AIRPORTS.SIN }}
		/>
		<FlightSparkRoute
			flight={{ originAirport: FIXTURE_AIRPORTS.SFO, destinationAirport: FIXTURE_AIRPORTS.SCL }}
		/>
		<FlightSparkRoute
			flight={{ originAirport: FIXTURE_AIRPORTS.SIN, destinationAirport: FIXTURE_AIRPORTS.EWR }}
		/>
		<FlightSparkRoute
			flight={{ originAirport: FIXTURE_AIRPORTS.LHR, destinationAirport: FIXTURE_AIRPORTS.AMS }}
		/>
	</div>
</section>
