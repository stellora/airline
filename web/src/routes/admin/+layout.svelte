<script lang="ts">
	import { page } from '$app/state'
	import { Badge } from '$lib/components/ui/badge'
	import UserCog from 'lucide-svelte/icons/user-cog'
	import AdminBreadcrumb from './admin-breadcrumb.svelte'

	let { children } = $props()
</script>

<Badge variant="secondary" class="self-start mb-3 inline-flex items-center gap-1"
	><UserCog size="2ch" /> Admin</Badge
>
<AdminBreadcrumb breadcrumbs={page.data.breadcrumbs} />
{@render children?.()}
