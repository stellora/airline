import type { Airport, Flight } from '$lib/types'
import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import FlightCard from './flight-card.svelte'

describe('FlightCard', () => {
	const mockFlight: Flight = {
		id: '1',
		title: 'Test Flight',
		starred: false
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightCard, { props: { flight: mockFlight } })
		expect(getByText('Test Flight')).toBeInTheDocument()
	})

	it('renders airports when provided', () => {
		const mockAirports: Airport[] = [
			{ id: '1', title: 'Airport 1' },
			{ id: '2', title: 'Airport 2' }
		]
		const { getByText } = render(FlightCard, {
			props: { flight: mockFlight, airports: mockAirports }
		})
		expect(getByText('Airport 1')).toBeInTheDocument()
		expect(getByText('Airport 2')).toBeInTheDocument()
	})
})
