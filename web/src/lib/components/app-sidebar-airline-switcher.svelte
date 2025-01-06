<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import * as Sidebar from '$lib/components/ui/sidebar/index.js'
	import { route } from '$lib/route-helpers'
	import type { Airline } from '$lib/types'
	import { cn } from '$lib/utils'
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down'
	import type { HTMLAttributes } from 'svelte/elements'
	import AirlineCode from './airline-code.svelte'

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
							<AirlineCode
								airline={activeAirline}
								tooltip={false}
								icon
								showName
								class="[&_[data-airline-name]]:text-xs"
							/>
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
								href={route('/admin/airlines/[airlineSpec]', {
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
