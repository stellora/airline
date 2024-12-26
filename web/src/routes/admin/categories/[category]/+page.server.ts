import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const category = (
		await apiClient.GET('/categories/{id}', { params: { path: { id: params.category } }, fetch })
	).data
	if (!category) {
		return fail(404)
	}
	const { productsInCategory, productsNotInCategory } = (
		await apiClient.GET('/categories/{categoryId}/products', {
			params: { path: { categoryId: params.category } },
			fetch
		})
	).data!
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

		const resp = await apiClient.PUT('/products/{productId}/categories/{categoryId}', {
			params: { path: { productId: product, categoryId: category } },
			body: { value },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text()
			})
		}
	},
	delete: async ({ request }) => {
		const data = await request.formData()
		const id = data.get('id')
		if (!id || typeof id !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		const resp = await apiClient.DELETE('/categories/{id}', {
			params: { path: { id } },
			fetch
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text()
			})
		}
		return redirect(303, '/admin/categories')
	}
}
