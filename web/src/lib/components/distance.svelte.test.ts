import { render } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import Distance from './distance.svelte'

describe('Distance', () => {
	it('renders nothing when distance is undefined', () => {
		const { container } = render(Distance, {})
		expect(container.textContent).toBe('')
	})

	it('renders singular mile when distance is 1', () => {
		const { container } = render(Distance, { props: { distanceMiles: 1 } })
		expect(container.textContent?.trim()).toBe('1 mile')
	})

	it('renders plural miles for other distances', () => {
		const { container } = render(Distance, { props: { distanceMiles: 5 } })
		expect(container.textContent?.trim()).toBe('5 miles')
	})

	it('rounds decimal distances', () => {
		const { container } = render(Distance, { props: { distanceMiles: 5.7 } })
		expect(container.textContent?.trim()).toBe('6 miles')
	})
})
