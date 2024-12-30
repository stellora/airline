import { apiClient } from '$lib/api'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load: PageServerLoad = async () => {
	return {
		airports: (await apiClient.GET('/airports', { fetch })).data,
	}
}

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData()
		const iataCode = data.get('iataCode')
		if (iataCode === null || typeof iataCode !== 'string') {
			return fail(400, {
				iataCode: iataCode,
				error: 'iataCode is required',
			})
		}

		const resp = await apiClient.POST('/airports', { body: { iataCode }, fetch })
		if (!resp.response.ok) {
			// TODO(sqs)
			return fail(422, {
				iataCode: iataCode,
				error: await resp.response.text(),
			})
		}
	},
}
