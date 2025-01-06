import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	return {
		aircraft: (
			await apiClient.GET('/airlines/{airlineSpec}/aircraft', {
				params: { path: { airlineSpec: params.airlineSpec } },
				fetch,
			})
		).data,
		form: await superValidate(
			{ airline: params.airlineSpec },
			typebox(schema['/aircraft']['POST']['args']['properties']['body']),
			{ errors: false },
		),
	}
}
