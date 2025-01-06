<script lang="ts">
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import BreadcrumbsForLayout from '$lib/components/ui/page/breadcrumbs-for-layout.svelte'
	import { route } from '$lib/route-helpers'
	import { cn } from '$lib/utils'

	const { children } = $props()
</script>

<BreadcrumbsForLayout>
	<Breadcrumb.Item>
		<Breadcrumb.Link href={route('/admin')}>
			{#snippet child({ props: { class: className, ...props } })}
				<a {...props} class={cn('admin-link', className)}>Global admin</a>
			{/snippet}
		</Breadcrumb.Link>
	</Breadcrumb.Item></BreadcrumbsForLayout
>

{@render children()}

<style>
	.admin-link {
		position: relative;
	}
	.admin-link::after {
		content: '';
		position: absolute;
		inset: -1rem -1.25rem -1rem -0.5rem;
		opacity: 0.15;
		background: repeating-linear-gradient(
			45deg,
			hsla(var(--warning)),
			hsla(var(--warning)) 10px,
			#000000 10px,
			#000000 20px
		);
	}
</style>
