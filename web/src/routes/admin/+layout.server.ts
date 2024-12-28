import { breadcrumbEntry } from '$lib/components/breadcrumbs'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async () => {
	return { ...(await breadcrumbEntry(null, 'Admin')) }
}
