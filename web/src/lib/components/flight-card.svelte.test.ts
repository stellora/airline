import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightCard from './flight-card.svelte'

describe('FlightCard', () => {
	const mockFlight: ComponentProps<typeof FlightCard>['flight'] = {
		id: 1,
		number: 'ST1',
		originAirport: { iataCode: 'AAA' },
		destinationAirport: { iataCode: 'BBB' },
		published: false,
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightCard, { props: { flight: mockFlight } })
		expect(getByText('ST1')).toBeInTheDocument()
	})
})
