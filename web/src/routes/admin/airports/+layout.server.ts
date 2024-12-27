import { breadcrumbEntry } from '../admin-breadcrumb'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ parent }) => {
	return {
		...(await breadcrumbEntry(parent, 'Airports'))
	}
}
