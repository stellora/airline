import type { Product } from '$lib/types'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ fetch }) => {
	return {
		productsWithCategories: await fetch('http://localhost:8080/products')
			.then((res) => res.json())
			.then((v) => {
				console.log(v)
				return v as Product[]
			})
	}
}
