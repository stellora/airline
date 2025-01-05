import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions } from './$types'

export const actions: Actions = {
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/itineraries/{itinerarySpec}', {
			params: { path: { itinerarySpec: params.itinerarySpec } },
			fetch,
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, '/admin/itineraries')
	},
}
