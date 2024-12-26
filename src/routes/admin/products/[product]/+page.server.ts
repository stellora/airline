import * as db from '$lib/server/database.js'
import { error, fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = ({ params }) => {
	const productWithCategories = db
		.listProductsWithCategories()
		.find(({ product }) => product.id === params.product)
	if (!productWithCategories) {
		throw error(404, 'Product not found')
	}
	return {
		productWithCategories
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
		db.deleteProduct(id)
		return redirect(303, '/admin/products')
	}
}
