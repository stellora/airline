<script module>
	import { FIXTURE_AIRPORTS } from '$lib/fixtures'
	import { defineMeta } from '@storybook/addon-svelte-csf'
	import GreatCircleRoute from './great-circle-route.svelte'

	const { Story } = defineMeta({
		title: 'maps/GreatCircleRoute',
		component: GreatCircleRoute,
		parameters: {
			layout: 'centered',
		},
		tags: ['autodocs'],
		args: {
			width: 'auto',
		},
	})
</script>

<Story name="One route" args={{ routes: [[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.EWR]] }} />

<Story
	name="Multiple routes"
	args={{
		routes: [
			[FIXTURE_AIRPORTS.SFO, FIXTURE_AIRPORTS.LHR],
			[FIXTURE_AIRPORTS.EWR, FIXTURE_AIRPORTS.LHR],
		],
	}}
/>
