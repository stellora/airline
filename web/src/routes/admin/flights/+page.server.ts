import { apiClient } from '$lib/api'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flights = (await apiClient.GET('/flights', { fetch })).data!
	return {
		flights
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

		const resp = await apiClient.POST('/flights', { body: { title }, fetch })
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				title,
				error: await resp.response.text()
			})
		}
	},

	setFlightStarred: async ({ request }) => {
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

		const resp = await apiClient.PATCH('/flights/{id}', {
			params: { path: { id } },
			body: { starred },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				starred: undefined,
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

		const resp = await apiClient.DELETE('/flights/{id}', {
			params: { path: { id } },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text()
			})
		}
	}
}
