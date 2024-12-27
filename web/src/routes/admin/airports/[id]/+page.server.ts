import { apiClient } from '$lib/api'
import { error, fail, redirect } from '@sveltejs/kit'
import { breadcrumbEntry } from '../../admin-breadcrumb'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params, parent }) => {
	const id = Number.parseInt(params.id)
	const airport = (await apiClient.GET('/airports/{id}', { params: { path: { id } }, fetch })).data
	if (!airport) {
		error(404)
	}
	const flights = apiClient
		.GET('/airports/{id}/flights', {
			params: { path: { id } },
			fetch
		})
		.then((resp) => resp.data)
	return {
		airport,
		flights,
		...(await breadcrumbEntry(parent, airport.iataCode)) // TODO!(sqS): make this a promise
	}
}

export const actions: Actions = {
	delete: async ({ request }) => {
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				error: 'id is required'
			})
		}
		const id = Number.parseInt(idStr)
		const resp = await apiClient.DELETE('/airports/{id}', {
			params: { path: { id } },
			fetch
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text()
			})
		}
		return redirect(303, '/admin/airports')
	}
}
