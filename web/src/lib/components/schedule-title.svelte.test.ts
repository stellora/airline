import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import ScheduleTitle from './schedule-title.svelte'

describe('ScheduleTitle', () => {
	const mockFlight: ComponentProps<typeof ScheduleTitle>['schedule'] = {
		id: 1,
		airline: { iataCode: 'XX', name: 'XX Airlines' },
		number: '1',
		originAirport: { iataCode: 'AAA', name: 'AAA Airport' },
		destinationAirport: { iataCode: 'BBB', name: 'BBB Airport' },
		published: false,
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(ScheduleTitle, { props: { schedule: mockFlight } })
		expect(getByText('XX 1')).toBeInTheDocument()
		expect(getByText('AAA–BBB')).toBeInTheDocument()
	})

	describe('publishing', () => {
		it('shows when flight is unpublished', () => {
			const { getByText } = render(ScheduleTitle, { props: { schedule: mockFlight } })
			expect(getByText('XX 1', { selector: 'span' })).toHaveClass('decoration-dotted')
		})
	})
})
