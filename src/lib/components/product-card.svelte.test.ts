import type { Product } from '$lib/types'
import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import ProductCard from './product-card.svelte'

describe('ProductCard', () => {
	const mockProduct: Product = {
		id: '1',
		title: 'Test Product',
		starred: false
	}

	it('renders product title correctly', () => {
		const { getByText } = render(ProductCard, { props: { product: mockProduct } })
		expect(getByText('Test Product')).toBeInTheDocument()
	})

	describe('starring', () => {
		it('shows star emoji when product is starred', () => {
			const starredProduct = { ...mockProduct, starred: true }
			const { getByText } = render(ProductCard, { props: { product: starredProduct } })
			expect(getByText('Test Product ⭐')).toBeInTheDocument()
		})
	})
})
