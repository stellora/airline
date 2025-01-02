import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const flightSchedules = apiClient
		.GET('/airlines/{airlineSpec}/flight-schedules', {
			params: { path: { airlineSpec: params.airlineSpec } },
			fetch,
		})
		.then((resp) => resp.data)
	return {
		flightSchedules,
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
