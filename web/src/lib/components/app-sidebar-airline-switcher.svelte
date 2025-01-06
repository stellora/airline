<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import * as Sidebar from '$lib/components/ui/sidebar/index.js'
	import { route } from '$lib/route-helpers'
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineCode from './airline-code.svelte'
	import AirlineIcon from './airline-icon.svelte'

	let {
		allAirlines,
		activeAirlineIataCode,
		class: className,
	}: {
		allAirlines: Airline[]
		activeAirlineIataCode: string | undefined
		class?: HTMLAttributes<never>['class']
	} = $props()

	let activeAirline = $derived(
		allAirlines.find((airline) => airline.iataCode === activeAirlineIataCode),
	)
</script>

<Sidebar.Menu>
	<Sidebar.MenuItem>
		<DropdownMenu.Root>
			<DropdownMenu.Trigger>
				{#snippet child({ props })}
					<Sidebar.MenuButton
						size="lg"
						class={cn(
							'data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground',
							className,
						)}
						{...props}
					>
						{#if activeAirline}
							<div class="flex items-center gap-1.5 overflow-hidden">
								<AirlineIcon
									airline={activeAirline}
									size="lg"
									showCode
									class="flex-shrink-0 mt-0.5 border border-foreground/25"
								/>
								<div class="flex flex-col gap-0.5 overflow-hidden">
									<span class="font-semibold text-sm truncate">{activeAirline.name}</span>
								</div>
							</div>
						{:else}
							<span class="text-muted-foreground">Airline...</span>
						{/if}
						<ChevronsUpDown class="ml-auto" />
					</Sidebar.MenuButton>
				{/snippet}
			</DropdownMenu.Trigger>
			<DropdownMenu.Content class="w-[--bits-dropdown-menu-anchor-width]" align="start">
				{#each allAirlines as airline (airline)}
					<DropdownMenu.Item>
						{#snippet child({ props })}
							<a
								{...props}
								href={route('/manage/[airlineSpec]', {
									params: { airlineSpec: airline.iataCode },
								})}
							>
								<AirlineCode
									{airline}
									tooltip={false}
									icon
									showName
									class={cn('[&_[data-airline-name]]:text-xs', {
										'font-bold': airline.iataCode === activeAirlineIataCode,
									})}
								/>
							</a>
						{/snippet}
					</DropdownMenu.Item>
				{/each}
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</Sidebar.MenuItem>
</Sidebar.Menu>
