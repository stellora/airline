import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ parent, params }) => {
	return {
		...(await breadcrumbEntry(parent, {
			url: `/admin/flight-schedules/${params.id}/manage`,
			title: 'Manage',
		})),
	}
}
