import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './airport-form'

export const load: PageServerLoad = async () => {
	return {
		airports: (await apiClient.GET('/airports', { fetch })).data,
		form: await superValidate(typebox(formSchema)),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/airports', {
			body: {
				iataCode: form.data.iataCode,
			},
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/admin/airports/[airportSpec]', {
				params: { airportSpec: resp.data.iataCode },
			}),
		)
	},
}
