import { Type } from '@sinclair/typebox'

export const formSchema = Type.Object({
	number: Type.String({ pattern: '^[A-Z0-9]{2}\\d{1,4}$' }),
	originAirport: Type.String({ minLength: 3, maxLength: 3 }),
	destinationAirport: Type.String({ minLength: 3, maxLength: 3 }),
	published: Type.Boolean(),
})

export type FormSchema = typeof formSchema
