import { env } from '$env/dynamic/public'
import type { paths } from '$lib/airline.openapi'
import createClient, { type Middleware } from 'openapi-fetch'

export const apiClient = createClient<paths>({
	baseUrl: `http://localhost:${env.PUBLIC_API_PORT ?? '8080'}`, // TODO(sqs)
})

const detectResponseError: Middleware = {
	async onResponse({ response }) {
		if (!response.ok) {
			// TODO!(sqs): does this work?
			throw new Error(`Error ${response.status}: ${response.statusText}`)
		}
		return response
	},
}
apiClient.use(detectResponseError)
