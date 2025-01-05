<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip'
	import type { SeatAssignment } from '$lib/types'

	const {
		rows = 30,
		seatsPerRow = 9,
		seatAssignments,
	}: {
		rows?: number
		seatsPerRow?: number
		seatAssignments: SeatAssignment[]
	} = $props()

	let bySeat = $derived(Object.fromEntries(seatAssignments.map((a) => [a.seat, a])))
</script>

<div class="flex flex-col gap-1 p-4">
	{#each Array(rows) as _, rowIndex}
		<div class="flex items-center gap-1">
			<div class="w-6 text-right text-sm text-muted-foreground">{rowIndex + 1}</div>
			{#each Array(seatsPerRow) as _, seatIndex}
				{@const seatLabel = `${rowIndex + 1}${String.fromCharCode(65 + seatIndex)}`}
				<Tooltip.Root
					disabled={!bySeat[seatLabel]}
					delayDuration={0}
					disableHoverableContent={true}
				>
					<Tooltip.Trigger>
						<button
							class="w-10 h-10 border border-border rounded bg-background hover:bg-muted disabled:bg-muted disabled:cursor-not-allowed text-sm"
							class:reserved={!bySeat[seatLabel]}
							disabled={!bySeat[seatLabel]}
						>
							{seatLabel}
						</button>
					</Tooltip.Trigger>
					<Tooltip.Content>
						{bySeat[seatLabel]?.passenger.name}
					</Tooltip.Content>
				</Tooltip.Root>
				{#if seatIndex !== 0 && seatIndex % 3 === 2}
					<div class="w-5"></div>
				{/if}
			{/each}
		</div>
	{/each}
</div>
