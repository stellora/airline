import { sveltekit } from '@sveltejs/kit/vite'
import { svelteTesting } from '@testing-library/svelte/vite'
import type { PluginOption } from 'vite'
import { defineConfig } from 'vitest/config'

export default defineConfig({
	plugins: [sveltekit() satisfies PluginOption as any, svelteTesting()], // eslint-disable-line @typescript-eslint/no-explicit-any
	test: {
		setupFiles: ['src/test/vitestSetup.ts'],
		include: ['src/**/*.test.ts'],
		environmentMatchGlobs: [['src/**/*.svelte.test.ts', 'jsdom']]
	}
})
