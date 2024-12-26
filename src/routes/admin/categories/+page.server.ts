import * as db from '$lib/server/database.js'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = () => {
	return {
		categories: db.listCategories()
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
			db.createCategory(title)
		} catch (error) {
			return fail(422, {
				title,
				error: error instanceof Error ? error.message : 'unknown error'
			})
		}
	}
}
