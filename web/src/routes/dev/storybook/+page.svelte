<script lang="ts">
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
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.LHR, FIXTURE_AIRPORTS.AMS]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.EWR, FIXTURE_AIRPORTS.SIN]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.SCL]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.LHR]]} />
		<GreatCircleRoute routes={[[FIXTURE_AIRPORTS.SYD, FIXTURE_AIRPORTS.EWR]]} />
		<GreatCircleRoute
			routes={[
				[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.SIN],
				[FIXTURE_AIRPORTS.LHR, FIXTURE_AIRPORTS.SCL],
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
	</div>
</section>
