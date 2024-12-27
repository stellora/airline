import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import AirportForm from './airport-form.svelte'

describe('AirportForm', () => {
	it('renders form elements', () => {
		render(AirportForm)
		expect(screen.getByPlaceholderText('IATA code')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add airport' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(AirportForm, { props: { form: { iataCode: '', error: 'Test error message' } } })
		expect(screen.getByText('Test error message', { exact: false })).toBeInTheDocument()
	})

	it('preserves input value from form data', () => {
		render(AirportForm, { props: { form: { iataCode: 'AAA', error: '' } } })
		expect(screen.getByPlaceholderText('IATA code')).toHaveValue('AAA')
	})

	it('requires title input', () => {
		render(AirportForm)
		const input = screen.getByPlaceholderText('IATA code')
		expect(input).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(AirportForm)
		const form = screen.queryByTestId('airport-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
