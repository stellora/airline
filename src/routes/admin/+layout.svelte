<script lang="ts">
	import { Badge } from '$lib/components/ui/badge'
	import UserCog from 'lucide-svelte/icons/user-cog'

	let { children } = $props()
</script>

<Badge variant="secondary" class="self-start mb-3 inline-flex items-center gap-1"
	><UserCog size="2ch" /> Admin</Badge
>
{@render children?.()}
