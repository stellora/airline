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
			typebox(schema['/flights/{id}']['PATCH']['args']['properties']['body']),
		)
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.PATCH('/flights/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/flights/[id]', { params: { id: resp.data.id.toString() } }),
		)
	},
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/flights/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			fetch,
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, route('/admin/flights'))
	},
}
