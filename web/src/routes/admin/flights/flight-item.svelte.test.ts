import type { Flight } from '$lib/types'
import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import FlightItem from './flight-item.svelte'

describe('FlightItem', () => {
	const mockFlight: Flight = {
		id: '1',
		title: 'Test Flight',
		published: false
	}

	it('renders flight title correctly', () => {
		const { getByText } = render(FlightItem, { props: { flight: mockFlight } })
		expect(getByText('Test Flight')).toBeInTheDocument()
	})

	describe('starring', () => {
		it('renders star/unstar button with correct text', async () => {
			const { getByText, rerender } = render(FlightItem, { props: { flight: mockFlight } })
			expect(getByText('Star')).toBeInTheDocument()

			await rerender({ flight: { ...mockFlight, published: true } })
			expect(getByText('Unstar')).toBeInTheDocument()
		})

		it('includes correct form data for starring/unstarring', () => {
			const { container } = render(FlightItem, { props: { flight: mockFlight } })
			const starForm = container.querySelector('form[action="?/setFlightPublished"]')
			const inputs = starForm?.querySelectorAll('input')

			expect(inputs?.[0]).toHaveValue('1')
			expect(inputs?.[1]).toHaveValue('true')
		})
	})

	describe('deletion', () => {
		it('includes correct form data for deletion', () => {
			const { container } = render(FlightItem, { props: { flight: mockFlight } })
			const deleteForm = container.querySelector('form[action="?/delete"]')
			const input = deleteForm?.querySelector('input')

			expect(input).toHaveValue('1')
		})
	})
})
