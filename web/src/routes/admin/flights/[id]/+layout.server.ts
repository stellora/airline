import { apiClient } from '$lib/api'
import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import { flightTitle } from '$lib/flight-helpers'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params, parent }) => {
	const id = Number.parseInt(params.id)
	const resp = await apiClient.GET('/flight-schedules/{id}', {
		params: { path: { id } },
		fetch,
	})
	if (!resp.response.ok || !resp.data) {
		// TODO(sqs)
		throw error(404, 'Flight not found')
	}
	const flight = resp.data
	return {
		flight,
		...(await breadcrumbEntry(parent, {
			url: `/admin/flight-schedules/${flight.id}`,
			title: flightTitle(flight),
		})),
	}
}
