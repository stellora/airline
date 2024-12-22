import * as db from '$lib/server/database.js';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = () => {
	return {
		products: db.getProducts()
	};
}

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData();

		try {
			db.createProduct(data.get('description'));
		} catch (error) {
			return fail(422, {
				description: data.get('description'),
				error: error instanceof Error ? error.message : 'Unknown error'
			});
		}
	},

	delete: async ({ request }) => {
		const data = await request.formData();
		db.deleteProduct(data.get('id'));
	}
};