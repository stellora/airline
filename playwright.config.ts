import { defineConfig } from '@playwright/test'

export default defineConfig({
	webServer: {
		command: 'npm run build && npm run preview',
		port: 4173
	},
	use: {
		browserName: 'chromium'
	},
	testDir: 'e2e'
})
