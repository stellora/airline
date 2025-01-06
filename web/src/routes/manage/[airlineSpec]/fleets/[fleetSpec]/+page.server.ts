import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import { route } from '$lib/route-helpers'
import { fail, redirect } from '@sveltejs/kit'
import { message, superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	return {
		fleetAircraft: await apiClient
			.GET('/airlines/{airlineSpec}/fleets/{fleetSpec}/aircraft', {
				params: {
					path: { airlineSpec: params.airlineSpec, fleetSpec: params.fleetSpec },
				},
				fetch,
			})
			.then((resp) => resp.data!),
	}
}

export const actions: Actions = {
	update: async ({ params, request }) => {
		const form = await superValidate(
			request,
			typebox(
				schema['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args']['properties']['body'],
			),
		)
		if (!form.valid) {
			return fail(400, { form })
		}
		const resp = await apiClient.PATCH('/airlines/{airlineSpec}/fleets/{fleetSpec}', {
			params: {
				path: { airlineSpec: params.airlineSpec, fleetSpec: params.fleetSpec },
			},
			body: form.data,
			fetch,
		})
		if (!resp.response.ok || !resp.data) {
			return message(form, resp.error, { status: 400 })
		}
		redirect(
			303,
			route('/manage/[airlineSpec]/fleets/[fleetSpec]', {
				params: {
					airlineSpec: resp.data.airline.iataCode,
					fleetSpec: resp.data.code,
				},
			}),
		)
	},
	delete: async ({ params }) => {
		const resp = await apiClient.DELETE('/airlines/{airlineSpec}/fleets/{fleetSpec}', {
			params: {
				path: { airlineSpec: params.airlineSpec, fleetSpec: params.fleetSpec },
			},
			fetch,
		})
		if (!resp.response.ok) {
			return fail(422, {
				error: await resp.response.text(),
			})
		}
		return redirect(
			303,
			route('/manage/[airlineSpec]/fleets', {
				params: { airlineSpec: params.airlineSpec },
			}),
		)
	},
}
