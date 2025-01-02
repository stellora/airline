import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightScheduleItem from './flight-schedule-item.svelte'

describe('FlightScheduleItem', () => {
	const mockFlightSchedule: ComponentProps<typeof FlightScheduleItem>['flight'] = {
		id: 1,
		airline: { iataCode: 'XX', name: 'XX Airlines' },
		number: '1',
		originAirport: { iataCode: 'AAA', name: 'AAA Airport' },
		destinationAirport: { iataCode: 'BBB', name: 'BBB Airport' },
		published: false,
	}

	it('renders flight number correctly', () => {
		const { getByText } = render(FlightScheduleItem, {
			props: { flightSchedule: mockFlightSchedule },
		})
		expect(getByText('XX1')).toBeInTheDocument()
	})
})
