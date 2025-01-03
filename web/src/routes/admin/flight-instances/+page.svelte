<script lang="ts">
	import { Button } from '$lib/components/ui/button'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import PageNav from '$lib/components/ui/page/page-nav.svelte'
	import Page from '$lib/components/ui/page/page.svelte'
	import Plus from 'lucide-svelte/icons/plus'
	import FlightScheduleForm from './[id]/flight-instance-form.svelte'
	import FlightScheduleList from './flight-instance-list.svelte'

	let { data } = $props()
</script>

<Page title="Flight instances">
	<FlightScheduleList flightInstances={data.flightInstances} />
</Page>
