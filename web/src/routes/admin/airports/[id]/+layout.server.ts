import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const id = Number.parseInt(params.id)
	const airport = (await apiClient.GET('/airports/{id}', { params: { path: { id } }, fetch })).data
	if (!airport) {
		error(404)
	}
	return {
		airport,
		...(await breadcrumbEntry(parent, {
			url: `/admin/airports/${airport.id}`,
			title: airport.iataCode,
		})),
	}
}
