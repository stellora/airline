import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flights = apiClient
		.GET('/airports/{airportSpec}/flights', {
			params: { path: { airportSpec: params.airportSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		flights,
	}
}

export const actions: Actions = {
	delete: async ({ request }) => {
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				error: 'id is required',
			})
		}
		const id = Number.parseInt(idStr)
		const resp = await apiClient.DELETE('/airports/{id}', {
			params: { path: { id } },
			fetch,
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, '/admin/airports')
	},
}
