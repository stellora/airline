<script lang="ts">
	import FlightForm from './flight-form.svelte'

	let { data } = $props()

	// TODO!(sqs): for testing +page.svelte files directly: https://github.com/sveltejs/kit/issues/5525#issuecomment-1186390654
</script>

<h1 class="text-4xl font-bold mb-4">New flight</h1>

<FlightForm form={data.form} />
