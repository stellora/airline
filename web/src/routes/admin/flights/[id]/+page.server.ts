import { apiClient } from '$lib/api'
import { error, fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/flights/{id}', {
		params: { path: { id: params.flight } },
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
		return redirect(303, '/admin/flights')
	}
}
