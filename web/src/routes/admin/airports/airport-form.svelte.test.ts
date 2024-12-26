import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import AirportForm from './airport-form.svelte'

describe('AirportForm', () => {
	it('renders form elements', () => {
		render(AirportForm)
		expect(screen.getByPlaceholderText('New airport...')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add airport' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(AirportForm, { props: { form: { title: '', error: 'Test error message' } } })
		expect(screen.getByText('Test error message', { exact: false })).toBeInTheDocument()
	})

	it('preserves input value from form data', () => {
		render(AirportForm, { props: { form: { title: 'Test Airport', error: '' } } })
		expect(screen.getByPlaceholderText('New airport...')).toHaveValue('Test Airport')
	})

	it('requires title input', () => {
		render(AirportForm)
		const input = screen.getByPlaceholderText('New airport...')
		expect(input).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(AirportForm)
		const form = screen.queryByTestId('airport-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
