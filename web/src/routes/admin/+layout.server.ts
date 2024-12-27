import type { LayoutServerLoad } from './$types'
import { breadcrumbEntry } from './admin-breadcrumb'

export const load: LayoutServerLoad = async () => {
	return { ...(await breadcrumbEntry(null, 'Admin')) }
}
