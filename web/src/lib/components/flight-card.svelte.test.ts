import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightCard from './flight-card.svelte'

describe('FlightCard', () => {
	const mockFlight: ComponentProps<typeof FlightCard>['schedule'] = {
		id: 1,
		airline: { iataCode: 'XX', name: 'XX Airlines' },
		number: '1',
		originAirport: { iataCode: 'AAA', name: 'AAA Airport' },
		destinationAirport: { iataCode: 'BBB', name: 'BBB Airport' },
		published: false,
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightCard, { props: { schedule: mockFlight } })
		expect(getByText('XX1')).toBeInTheDocument()
	})
})
