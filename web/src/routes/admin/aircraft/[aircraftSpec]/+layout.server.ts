import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import { existingAircraftToFormData, formSchema } from '../aircraft-form'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const resp = await apiClient.GET('/aircraft/{aircraftSpec}', {
		params: { path: { aircraftSpec: params.aircraftSpec } },
		fetch,
	})
	const aircraft = resp.data
	if (!aircraft) {
		error(resp.response.status, resp.error)
	}
	return {
		...(await breadcrumbEntry(parent, {
			url: route('/admin/aircraft/[aircraftSpec]', {
				params: { aircraftSpec: aircraft.registration },
			}),
			title: `${aircraft.registration} (${aircraft.airline.iataCode} ${aircraft.aircraftType})`,
		})),
		aircraft,
		form: await superValidate(existingAircraftToFormData(aircraft), typebox(formSchema)),
	}
}
