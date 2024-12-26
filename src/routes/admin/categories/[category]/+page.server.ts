import * as db from '$lib/server/database.js'
import { error, fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = ({ params }) => {
	const categories = db.listCategories()
	const category = categories.find((c) => c.id === params.category)

	if (!category) {
		throw error(404, 'Category not found')
	}

	return {
		category,
		productsInCategory: db.listProducts(),
		productsNotInCategory: db.listProducts()
	}
}
export const actions: Actions = {
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
