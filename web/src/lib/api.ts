import type { paths } from '$lib/airline.openapi'
import createClient from 'openapi-fetch'

export const apiClient = createClient<paths>({
	baseUrl: `http://localhost:${process.env.API_SERVER_PORT ?? '8080'}`, // TODO(sqs)
})
