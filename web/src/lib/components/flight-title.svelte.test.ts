import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightTitle from './flight-title.svelte'

describe('FlightTitle', () => {
	const mockFlight: ComponentProps<typeof FlightTitle>['flight'] = {
		id: 1,
		airline: { iataCode: 'XX', name: 'XX Airlines' },
		number: '1',
		originAirport: { iataCode: 'AAA', name: 'AAA Airport' },
		destinationAirport: { iataCode: 'BBB', name: 'BBB Airport' },
		published: false,
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightTitle, { props: { flight: mockFlight } })
		expect(getByText('XX1')).toBeInTheDocument()
		expect(getByText('AAA–BBB')).toBeInTheDocument()
	})

	describe('publishing', () => {
		it('shows when flight is unpublished', () => {
			const { getByText } = render(FlightTitle, { props: { flight: mockFlight } })
			expect(getByText('XX1', { selector: 'span' })).toHaveClass('decoration-dotted')
		})
	})
})
