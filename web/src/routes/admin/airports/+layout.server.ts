import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ parent }) => {
	return {
		...(await breadcrumbEntry(parent, { url: '/admin/airports', title: 'Airports' })),
	}
}
