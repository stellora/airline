import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions } from './$types'

export const actions: Actions = {
	update: async ({ params, request }) => {
		const form = await superValidate(
			request,
			typebox(schema['/aircraft/{aircraftSpec}']['PATCH']['args']['properties']['body']),
		)
		if (!form.valid) {
			return fail(400, { form })
		}
		const resp = await apiClient.PATCH('/aircraft/{aircraftSpec}', {
			params: { path: { aircraftSpec: params.aircraftSpec } },
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/manage/[airlineSpec]/aircraft/[aircraftSpec]', {
				params: { airlineSpec: params.airlineSpec, aircraftSpec: resp.data.registration },
			}),
		)
	},
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/aircraft/{aircraftSpec}', {
			params: { path: { aircraftSpec: params.aircraftSpec } },
			fetch,
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(
			303,
			route('/manage/[airlineSpec]/aircraft', { params: { airlineSpec: params.airlineSpec } }),
		)
	},
}
