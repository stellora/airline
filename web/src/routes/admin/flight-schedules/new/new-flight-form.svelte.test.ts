import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import NewFlightForm from './new-flight-form.svelte'

describe('NewFlightForm', () => {
	it('renders form elements', () => {
		render(NewFlightForm)
		expect(screen.getByPlaceholderText('Flight number')).toBeInTheDocument()
		expect(screen.getByPlaceholderText('From')).toBeInTheDocument()
		expect(screen.getByPlaceholderText('To')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add flight' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(NewFlightForm, { props: { form: { error: 'Test error message' } } })
		expect(screen.getByText('Test error message', { exact: false })).toBeInTheDocument()
	})

	it('preserves input values from form data', () => {
		render(NewFlightForm, {
			props: {
				form: {
					number: 'AB123',
					originAirport: 'LAX',
					destinationAirport: 'JFK',
					error: '',
				},
			},
		})
		expect(screen.getByPlaceholderText('Flight number')).toHaveValue('AB123')
		expect(screen.getByPlaceholderText('From')).toHaveValue('LAX')
		expect(screen.getByPlaceholderText('To')).toHaveValue('JFK')
	})

	it('requires all inputs', () => {
		render(NewFlightForm)
		expect(screen.getByPlaceholderText('Flight number')).toHaveAttribute('required')
		expect(screen.getByPlaceholderText('From')).toHaveAttribute('required')
		expect(screen.getByPlaceholderText('To')).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(NewFlightForm)
		const form = screen.getByTestId('flight-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
