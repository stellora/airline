import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flights = apiClient
		.GET('/airlines/{airlineSpec}/flights', {
			params: { path: { airlineSpec: params.airlineSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		flights,
	}
}

export const actions: Actions = {
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/airlines/{airlineSpec}', {
			params: { path: { airlineSpec: params.airlineSpec } },
			fetch,
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, '/admin/airlines')
	},
}
