import { apiClient } from '$lib/api'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	return {
		categories: (await apiClient.GET('/categories', { fetch })).data
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

		const resp = await apiClient.POST('/categories', { body: { title }, fetch })
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				title,
				error: await resp.response.text()
			})
		}
	}
}
