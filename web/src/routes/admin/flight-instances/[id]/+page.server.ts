import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './flight-instance-form'

export const load: PageServerLoad = async ({ params }) => {
	return {}
}

export const actions: Actions = {
	update: async ({ params, request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.PATCH('/flight-instances/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			body: {
				aircraft: form.data.aircraft,
				notes: form.data.notes,
			},
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/flight-instances/[id]', { params: { id: resp.data.id.toString() } }),
		)
	},
	delete: async ({ params, request }) => {
		const data = await request.formData()
		const resp = await apiClient.DELETE('/flight-instances/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			fetch,
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, route('/admin/flight-instances'))
	},
}
