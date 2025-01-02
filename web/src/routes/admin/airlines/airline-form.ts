import { Type } from '@sinclair/typebox'

// TODO!(sqs): use codegen
export const formSchema = Type.Object({
	iataCode: Type.String({ minLength: 2, maxLength: 2 }),
	name: Type.String({}),
})

export type FormSchema = typeof formSchema
