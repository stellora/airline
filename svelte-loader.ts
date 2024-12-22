import { GlobalRegistrator } from '@happy-dom/global-registrator'
import { plugin } from 'bun'
import { afterEach, beforeEach } from 'bun:test'
import { readFileSync } from 'fs'
import { compile } from 'svelte/compiler'

beforeEach(async () => {
	await GlobalRegistrator.register()
})

afterEach(async () => {
	await GlobalRegistrator.unregister()
})

plugin({
	name: 'svelte loader',
	setup(builder) {
		builder.module('$app/forms', () => {
			return {
				exports: {
					enhance: undefined
				},
				loader: 'object'
			}
		})

		builder.onLoad({ filter: /\.svelte(\?[^.]+)?$/ }, ({ path }) => {
			try {
				const source = readFileSync(
					path.substring(0, path.includes('?') ? path.indexOf('?') : path.length),
					'utf-8'
				)

				const result = compile(source, {
					filename: path,
					generate: 'client',
					dev: false
				})

				return {
					contents: result.js.code,
					loader: 'js'
				}
			} catch (err) {
				throw new Error(`Failed to compile Svelte component: ${err.message}`)
			}
		})
	}
})
