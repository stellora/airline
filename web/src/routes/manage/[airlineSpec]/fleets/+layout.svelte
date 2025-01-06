<script lang="ts">
	import { page } from '$app/state'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import { route } from '$lib/route-helpers'

	let { children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]/fleets', {
				params: { airlineSpec: page.params.airlineSpec },
			})}>Fleets</Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

{@render children()}
