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

const products: Product[] = INITIAL_PRODUCTS.map((title) => ({
	id: crypto.randomUUID(),
	title,
	starred: false
}))

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
		title,
		starred: false
	})
}

export function setProductStarred(id: string, starred: boolean): void {
	const product = products.find((product) => product.id === id)
	if (!product) {
		throw new Error(`product with id ${id} not found`)
	}
	product.starred = starred
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
