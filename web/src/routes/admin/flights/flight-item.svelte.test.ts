import { render } from '@testing-library/svelte'
import type { ComponentProps } from 'svelte'
import { describe, expect, it } from 'vitest'
import FlightItem from './flight-item.svelte'

describe('FlightItem', () => {
	const mockFlight: ComponentProps<typeof FlightItem>['flight'] = {
		id: 1,
		number: 'ST1',
		originAirport:{iataCode:'AAA'},
		destinationAirport: {iataCode:'BBB'},
		published: false
	}

	it('renders flight number correctly', () => {
		const { getByText } = render(FlightItem, { props: { flight: mockFlight } })
		expect(getByText('ST1')).toBeInTheDocument()
	})

	describe('publishing', () => {
		it('renders publish/unpublish button with correct text', async () => {
			const { getByText, rerender } = render(FlightItem, { props: { flight: mockFlight } })
			expect(getByText('Publish')).toBeInTheDocument()

			await rerender({ flight: { ...mockFlight, published: true } })
			expect(getByText('Unpublish')).toBeInTheDocument()
		})

		it('includes correct form data for publishing/unpublishing', () => {
			const { container } = render(FlightItem, { props: { flight: mockFlight } })
			const publishForm = container.querySelector('form[action="?/setFlightPublished"]')
			const inputs = publishForm?.querySelectorAll('input')

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
