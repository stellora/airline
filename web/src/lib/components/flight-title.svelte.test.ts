import type { Flight } from '$lib/types'
import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import FlightCard from './flight-card.svelte'

describe('FlightTitle', () => {
	const mockFlight: Flight = {
		id: '1',
		title: 'Test Flight',
		published: false
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightCard, { props: { flight: mockFlight } })
		expect(getByText('Test Flight')).toBeInTheDocument()
	})

	describe('starring', () => {
		it('shows star when flight is published', () => {
			const publishedFlight = { ...mockFlight, published: true }
			const { getByLabelText } = render(FlightCard, { props: { flight: publishedFlight } })
			expect(getByLabelText('Published')).toBeInTheDocument()
		})
	})
})
