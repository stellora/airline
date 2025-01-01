import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { route } from '$lib/route-helpers'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const resp = await apiClient.GET('/airlines/{airlineSpec}', {
		params: { path: { airlineSpec: params.airlineSpec } },
		fetch,
	})
	const airline = resp.data
	if (!airline) {
		error(resp.response.status, resp.error)
	}
	return {
		airline,
		...(await breadcrumbEntry(parent, {
			url: route('/admin/airlines/[airlineSpec]', {
				params: { airlineSpec: airline.iataCode },
			}),
			title: airline.iataCode,
		})),
	}
}
