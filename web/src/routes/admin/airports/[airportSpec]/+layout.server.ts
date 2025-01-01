import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const resp = await apiClient.GET('/airports/{airportSpec}', {
		params: { path: { airportSpec: params.airportSpec } },
		fetch,
	})
	const airport = resp.data
	console.log('XX', resp, params)
	if (!airport) {
		error(resp.response.status, resp.error)
	}
	return {
		airport: resp,
		...(await breadcrumbEntry(parent, {
			url: route('/admin/airports/[airportSpec]', {
				params: { airportSpec: airport.iataCode },
			}),
			title: airport.iataCode,
		})),
	}
}
