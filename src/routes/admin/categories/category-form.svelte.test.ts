import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import CategoryForm from './category-form.svelte'

describe('CategoryForm', () => {
	it('renders form elements', () => {
		render(CategoryForm)
		expect(screen.getByPlaceholderText('New category...')).toBeInTheDocument()
		expect(screen.getByRole('button', { name: 'Add category' })).toBeInTheDocument()
	})

	it('displays error message when form.error exists', () => {
		render(CategoryForm, { props: { form: { title: '', error: 'Test error message' } } })
		expect(screen.getByText('Test error message', { exact: false })).toBeInTheDocument()
	})

	it('preserves input value from form data', () => {
		render(CategoryForm, { props: { form: { title: 'Test Category', error: '' } } })
		expect(screen.getByPlaceholderText('New category...')).toHaveValue('Test Category')
	})

	it('requires title input', () => {
		render(CategoryForm)
		const input = screen.getByPlaceholderText('New category...')
		expect(input).toHaveAttribute('required')
	})

	it('has correct form action and method', () => {
		render(CategoryForm)
		const form = screen.queryByTestId('category-form')
		expect(form).toHaveAttribute('action', '?/create')
		expect(form).toHaveAttribute('method', 'POST')
	})
})
