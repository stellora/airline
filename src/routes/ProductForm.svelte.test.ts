import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import ProductForm from './ProductForm.svelte'

describe('ProductForm', () => {
	it('renders form elements', () => {
		render(ProductForm)
		expect(screen.getByPlaceholderText('New product...')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(ProductForm, { props: { form: { error: 'Test error message' } } })
		expect(screen.getByText('Test error message')).toBeInTheDocument()
	})

	it('preserves input value from form data', () => {
		render(ProductForm, { props: { form: { title: 'Test Product', error: '' } } })
		expect(screen.getByPlaceholderText('New product...')).toHaveValue('Test Product')
	})

	it('requires title input', () => {
		render(ProductForm)
		const input = screen.getByPlaceholderText('New product...')
		expect(input).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(ProductForm)
		const form = screen.queryByTestId('product-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
