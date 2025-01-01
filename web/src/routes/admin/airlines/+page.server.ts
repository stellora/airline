import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './airline-form'

export const load: PageServerLoad = async () => {
	return {
		airlines: (await apiClient.GET('/airlines', { fetch })).data,
		form: await superValidate(typebox(formSchema)),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/airlines', {
			body: {
				iataCode: form.data.iataCode,
				name: form.data.name,
			},
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
			route('/admin/airlines/[airlineSpec]', {
				params: { airlineSpec: resp.data.iataCode },
			}),
		)
	},
}
