import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { formSchema } from '../../flight-schedule-form'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params, parent }) => {
	const { flightSchedule } = await parent()
	return { form: await superValidate(flightSchedule, typebox(formSchema)) }
}

export const actions: Actions = {
	update: async ({ params, request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.PATCH('/flight-schedules/{id}', {
			params: { path: { id: Number.parseInt(params.id) } },
			body: {
				number: form.data.number,
				// TODO!(sqs): un-reuse the form across New and Update Flight
				//
				// TODO!(sqs)
				// originAirport: form.data.originAirport,
				// destinationAirport: form.data.destinationAirport,
				published: form.data.published,
			},
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/flight-schedules/[id]/manage', { params: { id: resp.data.id.toString() } }),
		)
	},
	setFlightSchedulePublished: async ({ request }) => {
		// TODO!(sqs): make this use the id from the URL not the form data
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				id: idStr,
				error: 'id is required',
			})
		}
		const id = Number.parseInt(idStr)

		const publishedStr = data.get('published')
		if (publishedStr !== 'true' && publishedStr !== 'false') {
			return fail(400, {
				published: undefined,
				error: 'published must be "true" or "false"',
			})
		}
		const published = publishedStr === 'true'

		const resp = await apiClient.PATCH('/flight-schedules/{id}', {
			params: { path: { id } },
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
	delete: async ({ request }) => {
		// TODO!(sqs): make this use the id from the URL not the form data
		const data = await request.formData()
		const idStr = data.get('id')
		if (!idStr || typeof idStr !== 'string') {
			return fail(400, {
				error: 'id is required',
			})
		}
		const id = Number.parseInt(idStr)

		const resp = await apiClient.DELETE('/flight-schedules/{id}', {
			params: { path: { id } },
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
