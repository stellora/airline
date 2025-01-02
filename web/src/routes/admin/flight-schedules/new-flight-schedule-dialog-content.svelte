<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import type { ComponentProps } from 'svelte'
	import FlightScheduleForm from './flight-schedule-form.svelte'

	const props: ComponentProps<typeof FlightScheduleForm> = $props()
</script>

<Dialog.Content>
	<Dialog.Header>
		<Dialog.Title>New flight schedule</Dialog.Title>
	</Dialog.Header>
	<FlightScheduleForm {...props} />
</Dialog.Content>
