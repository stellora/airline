import type { Product, ProductCategory } from '$lib/types'
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

	it('renders categories when provided', () => {
		const mockCategories: ProductCategory[] = [
			{ id: '1', title: 'Category 1' },
			{ id: '2', title: 'Category 2' }
		]
		const { getByText } = render(ProductCard, {
			props: { product: mockProduct, categories: mockCategories }
		})
		expect(getByText('Category 1')).toBeInTheDocument()
		expect(getByText('Category 2')).toBeInTheDocument()
	})
})
