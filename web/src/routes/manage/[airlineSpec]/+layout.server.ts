import { apiClient } from '$lib/api'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async ({ params }) => {
	const resp = await apiClient.GET('/airlines/{airlineSpec}', {
		params: { path: { airlineSpec: params.airlineSpec } },
		fetch,
	})
	const airline = resp.data
	if (!airline) {
		error(resp.response.status, resp.error)
	}
	return {
		airline,
	}
}
