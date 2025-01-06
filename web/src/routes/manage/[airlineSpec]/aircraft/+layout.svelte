<script lang="ts">
	import { page } from '$app/state'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import { route } from '$lib/route-helpers'

	const { children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link
			href={route('/manage/[airlineSpec]/aircraft', {
				params: { airlineSpec: page.params.airlineSpec },
			})}>Aircraft</Breadcrumb.Link
		>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

{@render children()}
