import { env } from '$env/dynamic/public'
import type { paths } from '$lib/airline.openapi'
import createClient from 'openapi-fetch'

export const apiClient = createClient<paths>({
	baseUrl: `http://localhost:${env.PUBLIC_API_PORT ?? '8080'}`, // TODO(sqs)
})
