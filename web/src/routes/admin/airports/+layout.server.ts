import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = ({ parent }) => {
	return { breadcrumbs: [parent().then((parent) => parent.breadcrumbs), 'Airports'] }
}
