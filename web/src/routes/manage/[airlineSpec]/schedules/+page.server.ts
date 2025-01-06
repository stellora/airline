import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	return {
		schedules: await apiClient
			.GET('/airlines/{airlineSpec}/schedules', {
				params: { path: { airlineSpec: params.airlineSpec } },
				fetch,
			})
			.then((resp) => resp.data!),
		form: await superValidate(
			{ airline: params.airlineSpec },
			typebox(schema['/schedules']['POST']['args']['properties']['body']),
			{ errors: false },
		),
	}
}

export const actions: Actions = {
	create: async ({ params, request }) => {
		const form = await superValidate(
			request,
			typebox(schema['/schedules']['POST']['args']['properties']['body']),
		)
		if (!form.valid) {
			return fail(400, { form })
		}

		const resp = await apiClient.POST('/schedules', {
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/manage/[airlineSpec]/schedules/[id]', {
				params: { airlineSpec: params.airlineSpec, id: resp.data.id.toString() },
			}),
		)
	},
}
