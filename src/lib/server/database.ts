import type { Product } from '$lib/types'

const INITIAL_PRODUCTS = [
	'Fork',
	'Spoon',
	'Knife',
	'Cast-Iron Pan',
	'Baking Sheet',
	'Flour',
	'Tomato',
	'Zucchini',
	'Avocado'
]

const products: Product[] = INITIAL_PRODUCTS.map((title) => ({ id: crypto.randomUUID(), title }))

export function listProducts() {
	return products
}

export function createProduct(title: string): void {
	if (title === '') {
		throw new Error('title must not be empty')
	}

	if (products.find((product) => product.title === title)) {
		throw new Error('title must be unique across all products')
	}

	products.push({
		id: crypto.randomUUID(),
		title: title
	})
}

export function deleteProduct(id: string): void {
	const index = products.findIndex((product) => product.id === id)
	if (index !== -1) {
		products.splice(index, 1)
	}
}

export function deleteAllProducts(): void {
	products.length = 0
}
