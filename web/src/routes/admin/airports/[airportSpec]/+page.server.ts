import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { formDataToAirportRequest, formSchema } from '../airport-form'
import type { Actions } from './$types'

export const actions: Actions = {
	update: async ({ params, request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}
		const resp = await apiClient.PATCH('/airports/{airportSpec}', {
			params: { path: { airportSpec: params.airportSpec } },
			body: formDataToAirportRequest(form.data),
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/airports/[airportSpec]', { params: { airportSpec: resp.data.iataCode } }),
		)
	},
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
