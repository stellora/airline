import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'
import { formSchema } from './flight-instance-form'

export const load: PageServerLoad = async () => {
	const flightInstances = await apiClient
		.GET('/flight-instances', { fetch })
		.then((resp) => resp.data!)
	return {
		flightInstances,
		form: await superValidate(typebox(formSchema)),
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const form = await superValidate(request, typebox(formSchema))
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/flight-instances', {
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
}
