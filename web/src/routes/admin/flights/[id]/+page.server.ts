import { apiClient } from '$lib/api'
import { error, fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flights/{id}', {
		params: { path: { id } },
		fetch
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Flight not found')
	}
	return {
		flight: resp.data
	}
}

export const actions: Actions = {
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
		return redirect(303, '/admin/flights')
	}
}
