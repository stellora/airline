import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightItem from './flight-item.svelte'

describe('FlightItem', () => {
	const mockFlight: ComponentProps<typeof FlightItem>['flight'] = {
		id: 1,
		number: 'ST1',
		originAirport: { iataCode: 'AAA' },
		destinationAirport: { iataCode: 'BBB' },
		published: false,
	}

	it('renders flight number correctly', () => {
		const { getByText } = render(FlightItem, { props: { flight: mockFlight } })
		expect(getByText('ST1')).toBeInTheDocument()
	})
})
