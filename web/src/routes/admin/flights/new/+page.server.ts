import { apiClient } from '$lib/api'
import { fail, redirect } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './flight-form'

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(typebox(formSchema)),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
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
			return fail(422, {
				form,
				error: 'Error creating flight',
			})
		}
		redirect(303, `/admin/flights/${resp.data.id}`)
	},
}
