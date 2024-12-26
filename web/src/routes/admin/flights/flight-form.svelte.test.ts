import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import FlightForm from './flight-form.svelte'

describe('FlightForm', () => {
	it('renders form elements', () => {
		render(FlightForm)
		expect(screen.getByPlaceholderText('New flight...')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add flight' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(FlightForm, { props: { form: { error: 'Test error message' } } })
		expect(screen.getByText('Test error message', { exact: false })).toBeInTheDocument()
	})

	it('preserves input value from form data', () => {
		render(FlightForm, { props: { form: { title: 'Test Flight', error: '' } } })
		expect(screen.getByPlaceholderText('New flight...')).toHaveValue('Test Flight')
	})

	it('requires title input', () => {
		render(FlightForm)
		const input = screen.getByPlaceholderText('New flight...')
		expect(input).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(FlightForm)
		const form = screen.queryByTestId('flight-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
