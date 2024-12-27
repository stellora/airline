import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightTitle from './flight-title.svelte'

describe('FlightTitle', () => {
	const mockFlight: ComponentProps<typeof FlightTitle>['flight'] = {
		id: 1,
		number: 'ST1',
		originAirport: { iataCode: 'AAA' },
		destinationAirport: { iataCode: 'BBB' },
		published: false
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightTitle, { props: { flight: mockFlight } })
		expect(getByText('ST1')).toBeInTheDocument()
		expect(getByText('AAA–BBB')).toBeInTheDocument()
	})

	describe('publishing', () => {
		it('shows when flight is unpublished', () => {
			const { getByText } = render(FlightTitle, { props: { flight: mockFlight } })
			expect(getByText('ST1', { selector: 'span' })).toHaveClass('decoration-dotted')
		})
	})
})
