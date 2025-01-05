import { env } from '$env/dynamic/public'
import type { paths } from '$lib/airline.openapi'
import createClient, { type Middleware } from 'openapi-fetch'
import { schema } from './airline.typebox'

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

// Workaround for the sveltekit-superforms error "Multi-type unions must have a default value, or exactly one of the union types must have.".
// components['schemas']['AirlineSpec'].default = undefined
// components['schemas']['AircraftSpec'].default = undefined
export function workaroundForMultiTypeUnions(): void {
	schema['/aircraft'].POST.args.properties.body.properties.airline.default = undefined
	schema['/aircraft/{aircraftSpec}'].PATCH.args.properties.body.properties.airline.default =
		undefined
}
