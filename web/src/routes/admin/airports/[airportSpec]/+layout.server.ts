import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const airport = await apiClient.GET('/airports/{airportSpec}', {
		params: { path: { airportSpec: params.airportSpec } },
		fetch,
	})
	console.log('XX', airport, params)
	if (!airport) {
		error(404)
	}
	return {
		airport,
		...(await breadcrumbEntry(parent, {
			url: route('/admin/airports/[airportSpec]', {
				params: { airportSpec: airport.iataCode },
			}),
			title: airport.iataCode,
		})),
	}
}
