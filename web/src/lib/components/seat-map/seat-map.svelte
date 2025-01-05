<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import type { SeatAssignment } from '$lib/types'

	const {
		seatConfig = [
			{
				name: 'First Class',
				rows: 2,
				seatsPerRow: 3,
				seatSize: 'w-24 h-28',
				aisles: [1, 2],
				rowSpacing: 'gap-2',
			},
			{
				name: 'Business Class',
				rows: 12,
				seatsPerRow: 4,
				seatSize: 'w-20 h-20',
				aisles: [1, 3],
				rowSpacing: 'gap-2',
			},
			{
				name: 'Premium Economy Plus',
				rows: 4,
				seatsPerRow: 6,
				seatSize: 'w-16 h-18',
				aisles: [2, 4],
				rowSpacing: 'gap-2',
			},
			{
				name: 'Economy Plus',
				rows: 7,
				seatsPerRow: 10,
				seatSize: 'w-14 h-14',
				aisles: [3, 7],
				rowSpacing: 'gap-2',
			},
			{
				name: 'Economy',
				rows: 10,
				seatsPerRow: 10,
				seatSize: 'w-14 h-14',
				aisles: [3, 7],
				rowSpacing: 'gap-1',
			},
		],
		seatAssignments,
	}: {
		seatConfig?: Array<{
			name: string
			rows: number
			seatsPerRow: number
			seatSize: string
			aisles: number[]
			rowSpacing: string
		}>
		seatAssignments: SeatAssignment[]
	} = $props()

	let bySeat = $derived(Object.fromEntries(seatAssignments.map((a) => [a.seat, a])))
	let rowOffsets = $derived(
		seatConfig.map((_, index) =>
			seatConfig.slice(0, index).reduce((sum, section) => sum + section.rows, 0),
		),
	)
</script>

<div class="flex flex-col gap-6 p-4">
	{#each seatConfig as section, sectionIndex}
		<div class="flex flex-col {section.rowSpacing}">
			{#each Array(section.rows) as _, rowIndex}
				<div class="flex items-center gap-1 justify-center">
					<div class="w-6 text-right text-sm text-muted-foreground">
						{rowIndex + rowOffsets[sectionIndex] + 1}
					</div>
					{#each Array(section.seatsPerRow) as _, seatIndex}
						{@const seatLabel = `${rowIndex + rowOffsets[sectionIndex] + 1}${String.fromCharCode(65 + seatIndex)}`}
						<Tooltip.Root
							disabled={!bySeat[seatLabel]}
							delayDuration={0}
							disableHoverableContent={true}
						>
							<Tooltip.Trigger>
								<button
									class="{section.seatSize} border border-border rounded bg-background hover:bg-muted disabled:bg-muted disabled:cursor-not-allowed text-sm overflow-hidden"
									class:reserved={!bySeat[seatLabel]}
									disabled={!bySeat[seatLabel]}
								>
									{seatLabel}
									{#if bySeat[seatLabel]}
										<span class="hidden lg:block text-muted-foreground text-xxs">
											{bySeat[seatLabel]?.passenger.name}
										</span>
									{/if}
								</button>
							</Tooltip.Trigger>
							<Tooltip.Content>
								{bySeat[seatLabel]?.passenger.name}
							</Tooltip.Content>
						</Tooltip.Root>
						{#if section.aisles.includes(seatIndex + 1)}
							<div class="w-5"></div>
						{/if}
					{/each}
				</div>
			{/each}
		</div>
	{/each}
</div>
