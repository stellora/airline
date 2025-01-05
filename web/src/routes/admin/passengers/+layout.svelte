<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import { route } from '$lib/route-helpers'

	const { children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link href={route('/admin/passengers')}>Passengers</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

{@render children()}
