import type { Preview } from '@storybook/svelte'
import { themes, type ThemeVars } from '@storybook/theming'
import '../src/app.css'
import './preview.css'

const preview: Preview & { docs?: { theme: ThemeVars } } = {
	parameters: {
		controls: {
			matchers: {
				color: /(background|color)$/i,
				date: /Date$/i,
			},
		},
	},
	docs: {
		theme: themes.dark,
	},
}
export default preview
