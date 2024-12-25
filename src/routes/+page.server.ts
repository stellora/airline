import * as db from '$lib/server/database.js'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = () => {
	return {
		products: db.listProducts()
	}
}
