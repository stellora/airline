import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './new-flight-form'

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(typebox(formSchema)),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		console.log('QQQ')
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/flights', {
			body: {
				number: form.data.number,
				originAirport: form.data.originAirport,
				destinationAirport: form.data.destinationAirport,
				published: false,
			},
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			// TODO(sqs)
			console.log('EEE', resp.error)
			return message(form, resp.error, { status: 400 })
		}
		redirect(303, `/admin/flights/${resp.data.id}`)
	},
}
