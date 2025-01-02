import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flights = apiClient
		.GET('/airports/{airportSpec}/flight-schedules', {
			params: { path: { airportSpec: params.airportSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		flights,
	}
}

export const actions: Actions = {
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/airports/{airportSpec}', {
			params: { path: { airportSpec: params.airportSpec } },
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
