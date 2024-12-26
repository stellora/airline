import * as db from '$lib/server/database.js'
import { error, fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = ({ params }) => {
	const categories = db.listCategories()
	const category = categories.find((c) => c.id === params.category)

	if (!category) {
		throw error(404, 'Category not found')
	}

	const { productsInCategory, productsNotInCategory } = db.listCategoryProducts(category.id)
	return {
		category,
		productsInCategory,
		productsNotInCategory
	}
}
export const actions: Actions = {
	setProductInCategory: async ({ request }) => {
		const data = await request.formData()

		const category = data.get('category')
		const product = data.get('product')
		if (!category || typeof category !== 'string') {
			return fail(400, { error: 'category is required' })
		}
		if (!product || typeof product !== 'string') {
			return fail(400, { error: 'product is required' })
		}

		const valueStr = data.get('value')
		if (valueStr !== 'true' && valueStr !== 'false') {
			return fail(400, {
				value: false,
				error: 'value must be "true" or "false"'
			})
		}
		const value = valueStr === 'true'

		db.setProductInCategory(product, category, value)
	},
	delete: async ({ request }) => {
		const data = await request.formData()
		const id = data.get('id')
		if (!id || typeof id !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		db.deleteCategory(id)
		return redirect(303, '/admin/categories')
	}
}
