import type { paths } from '$lib/shop.openapi'
import createClient from 'openapi-fetch'

export const apiClient = createClient<paths>({
	baseUrl: 'http://localhost:8080' // TODO(sqs)
})
