import * as db from '$lib/server/database.js'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = () => {
	return {
		products: db.listProducts()
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData()

		const title = data.get('title')
		if (title === null || typeof title !== 'string') {
			return fail(400, {
				title,
				error: 'title is required'
			})
		}
		try {
			db.createProduct(title)
		} catch (error) {
			return fail(422, {
				title,
				error: error instanceof Error ? error.message : 'unknown error'
			})
		}
	},

	setProductStarred: async ({ request }) => {
		const data = await request.formData()
		const id = data.get('id')
		if (!id || typeof id !== 'string') {
			return fail(400, {
				id,
				error: 'id is required'
			})
		}
		const starredStr = data.get('starred')
		if (starredStr !== 'true' && starredStr !== 'false') {
			return fail(400, {
				starred: undefined,
				error: 'starred must be "true" or "false"'
			})
		}
		const starred = starredStr === 'true'
		db.setProductStarred(id, starred)
	},

	delete: async ({ request }) => {
		const data = await request.formData()
		const id = data.get('id')
		if (!id || typeof id !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		db.deleteProduct(id)
	}
}
