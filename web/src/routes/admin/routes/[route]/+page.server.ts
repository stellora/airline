import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { flightRoute } from '$lib/flight-helpers'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ parent, params, url }) => {
	const route = await apiClient
		.GET('/routes/{route}', { fetch, params: { path: { route: params.route } } })
		.then((resp) => resp.data!)
	return {
		route,
		...breadcrumbEntry(parent, {
			url: url.pathname,
			title: flightRoute(route.originAirport, route.destinationAirport)
		})
	}
}
