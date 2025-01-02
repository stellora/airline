import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'
import { route } from '$lib/route-helpers'

export const load: PageServerLoad = async ({ params }) => {
	return {}
}

export const actions: Actions = {
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
	delete: async ({ params, request }) => {
		const data = await request.formData()
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
