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
				airline: form.data.airline,
				number: form.data.number,
				originAirport: form.data.originAirport,
				destinationAirport: form.data.destinationAirport,
				published: form.data.published,
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
	setFlightInstancePublished: async ({ request, params }) => {
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

		const resp = await apiClient.PATCH('/flight-instances/{id}', {
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
