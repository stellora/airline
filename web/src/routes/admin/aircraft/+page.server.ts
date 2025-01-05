import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	return {
		aircraft: (await apiClient.GET('/aircraft', { fetch })).data,
		form: await superValidate(typebox(schema['/aircraft']['POST']['args']['properties']['body'])),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const form = await superValidate(
			request,
			typebox(schema['/aircraft']['POST']['args']['properties']['body']),
		)
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/aircraft', {
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			// TODO!(sqs): if submitting the form has an error, the message is only shown
			// the first time you submit. if you click submit again, it does not show the
			// message.
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/aircraft/[aircraftSpec]', {
				params: { aircraftSpec: resp.data.registration },
			}),
		)
	},
}
