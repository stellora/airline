import { schema } from '$lib/airline.typebox'
import { apiClient } from '$lib/api'
import type { Fleet } from '$lib/types'
import type { Static } from '@sinclair/typebox'
import { superValidate } from 'sveltekit-superforms'
import { typebox } from 'sveltekit-superforms/adapters'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ params }) => {
	const fleet = await apiClient
		.GET('/airlines/{airlineSpec}/fleets/{fleetSpec}', {
			params: {
				path: { airlineSpec: params.airlineSpec, fleetSpec: params.fleetSpec },
			},
			fetch,
		})
		.then((resp) => resp.data!)
	return {
		fleet,
		form: await superValidate(
			existingFleetToFormData(fleet),
			typebox(
				schema['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args']['properties']['body'],
			),
		),
	}
}

function existingFleetToFormData(
	a: Fleet,
): Static<
	(typeof schema)['/airlines/{airlineSpec}/fleets/{fleetSpec}']['PATCH']['args']['properties']['body']
> {
	return {
		code: a.code,
		description: a.description,
	}
}
