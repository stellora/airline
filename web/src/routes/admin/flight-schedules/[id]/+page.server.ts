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
			typebox(schema['/flight-schedules/{id}']['PATCH']['args']['properties']['body']),
		)
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.PATCH('/flight-schedules/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
	},
	setFlightSchedulePublished: async ({ request, params }) => {
		// TODO!(sqs): make this use the id from the URL not the form data
		const data = await request.formData()

		const publishedStr = data.get('published')
		if (publishedStr !== 'true' && publishedStr !== 'false') {
			return fail(400, {
				published: undefined,
				error: 'published must be "true" or "false"',
			})
		}
		const published = publishedStr === 'true'

		const resp = await apiClient.PATCH('/flight-schedules/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			body: { published },
			fetch,
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				published: undefined,
				error: await resp.response.text(),
			})
		}
	},
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/flight-schedules/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			fetch,
		})
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(303, route('/admin/flight-schedules'))
	},
}
