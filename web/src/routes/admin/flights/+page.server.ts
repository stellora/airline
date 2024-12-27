import { apiClient } from '$lib/api'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	const flights = apiClient.GET('/flights', { fetch }).then((resp) => resp.data)
	return {
		flights: await flights
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData()
		const number = data.get('number')
		if (number === null || typeof number !== 'string') {
			return fail(400, {
				number,
				error: 'flight number is required'
			})
		}

		const resp = await apiClient.POST('/flights', { body: { number }, fetch })
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				number,
				error: await resp.response.text()
			})
		}
	},

	setFlightPublished: async ({ request }) => {
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				id: idStr,
				error: 'id is required'
			})
		}
		const id = Number.parseInt(idStr)

		const publishedStr = data.get('published')
		if (publishedStr !== 'true' && publishedStr !== 'false') {
			return fail(400, {
				published: undefined,
				error: 'published must be "true" or "false"'
			})
		}
		const published = publishedStr === 'true'

		const resp = await apiClient.PATCH('/flights/{id}', {
			params: { path: { id } },
			body: { published },
			fetch
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				published: undefined,
				error: await resp.response.text()
			})
		}
	},

	delete: async ({ request }) => {
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		const id = Number.parseInt(idStr)

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
