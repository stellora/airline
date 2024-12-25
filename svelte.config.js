import adapter from '@sveltejs/adapter-node'
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter()
	},

	compilerOptions: {
		// TODO(sqs): can set to `true` when https://github.com/lucide-icons/lucide/issues/2312 is
		// fixed.
		runes: undefined
	},
}

export default config
