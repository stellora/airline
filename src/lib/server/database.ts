import type { Product, ProductCategory } from '$lib/types'

const INITIAL_CATEGORIES = ['Silverware', 'Cookware', 'Vegetables']

const categories: ProductCategory[] = INITIAL_CATEGORIES.map((title) => ({
	id: crypto.randomUUID(),
	title
}))

export function listCategories(): ProductCategory[] {
	return categories
}

export function createCategory(title: string): void {
	if (title === '') {
		throw new Error('title must not be empty')
	}

	if (categories.find((category) => category.title === title)) {
		throw new Error('title must be unique across all categories')
	}

	categories.push({
		id: crypto.randomUUID(),
		title
	})
}

export function deleteCategory(id: string): void {
	const index = categories.findIndex((category) => category.id === id)
	if (index !== -1) {
		categories.splice(index, 1)
	}
}

const INITIAL_PRODUCTS = [
	'Fork',
	'Spoon',
	'Knife',
	'Cast-Iron Pan',
	'Baking Sheet',
	'Cutting Board',
	'Tomato',
	'Zucchini',
	'Avocado'
]

const products: Product[] = INITIAL_PRODUCTS.map((title) => ({
	id: crypto.randomUUID(),
	title,
	starred: false
}))

export function listProducts(): Product[] {
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
	productCategoryMemberships.length = 0
}

const productCategoryMemberships: { product: string; category: string }[] = []

export function setProductInCategory(product: string, category: string, value: boolean): void {
	const productObj = products.find((p) => p.id === product)
	if (!productObj) {
		throw new Error(`product with id ${product} not found`)
	}

	const categoryObj = categories.find((c) => c.id === category)
	if (!categoryObj) {
		throw new Error(`category with id ${category} not found`)
	}

	const existsIndex = productCategoryMemberships.findIndex(
		(m) => m.product === product && m.category === category
	)
	if (value) {
		if (existsIndex === -1) {
			productCategoryMemberships.push({ product, category })
		}
	} else {
		if (existsIndex !== -1) {
			productCategoryMemberships.splice(existsIndex, 1)
		}
	}
}

export function listCategoryProducts(category: string): {
	productsInCategory: Product[]
	productsNotInCategory: Product[]
} {
	const productsInCategory = products.filter((product) =>
		productCategoryMemberships.some(
			(membership) => membership.product === product.id && membership.category === category
		)
	)

	const productsNotInCategory = products.filter(
		(product) =>
			!productCategoryMemberships.some(
				(membership) => membership.product === product.id && membership.category === category
			)
	)

	return {
		productsInCategory,
		productsNotInCategory
	}
}

export function listProductsWithCategories(): {
	product: Product
	categories: ProductCategory[]
}[] {
	return products.map((product) => ({
		product,
		categories: categories.filter((category) =>
			productCategoryMemberships.some(
				(membership) => membership.product === product.id && membership.category === category.id
			)
		)
	}))
}
