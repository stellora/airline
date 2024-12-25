import type { Product } from '$lib/types'
import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import ProductItem from './product-item.svelte'

describe('ProductItem', () => {
	const mockProduct: Product = {
		id: '1',
		title: 'Test Product',
		starred: false
	}

	it('renders product title correctly', () => {
		const { getByText } = render(ProductItem, { props: { product: mockProduct } })
		expect(getByText('Test Product')).toBeInTheDocument()
	})

	describe('starring', () => {
		it('shows star emoji when product is starred', () => {
			const starredProduct = { ...mockProduct, starred: true }
			const { getByText } = render(ProductItem, { props: { product: starredProduct } })
			expect(getByText('Test Product ⭐')).toBeInTheDocument()
		})

		it('renders star/unstar button with correct text', async () => {
			const { getByText, rerender } = render(ProductItem, { props: { product: mockProduct } })
			expect(getByText('Star')).toBeInTheDocument()

			await rerender({ product: { ...mockProduct, starred: true } })
			expect(getByText('Unstar')).toBeInTheDocument()
		})

		it('includes correct form data for starring/unstarring', () => {
			const { container } = render(ProductItem, { props: { product: mockProduct } })
			const starForm = container.querySelector('form[action="?/setProductStarred"]')
			const inputs = starForm?.querySelectorAll('input')

			expect(inputs?.[0]).toHaveValue('1')
			expect(inputs?.[1]).toHaveValue('true')
		})
	})

	describe('deletion', () => {
		it('includes correct form data for deletion', () => {
			const { container } = render(ProductItem, { props: { product: mockProduct } })
			const deleteForm = container.querySelector('form[action="?/delete"]')
			const input = deleteForm?.querySelector('input')

			expect(input).toHaveValue('1')
		})
	})
})
