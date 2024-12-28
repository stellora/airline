import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import Breadcrumbs from './breadcrumbs.svelte'

describe('Breadcrumbs', () => {
	it('renders breadcrumbs correctly', async () => {
		render(Breadcrumbs, { props: { entries: ['A', 'B', 'C'] } })
		expect(readBreadcrumbs()).toStrictEqual('A>B>C')
	})

	it('renders loading state for async breadcrumbs', () => {
		render(Breadcrumbs, {
			props: { entries: [Promise.resolve('A'), Promise.resolve('B')] }
		})
		expect(screen.getAllByTestId('breadcrumb-ellipsis')).toHaveLength(2)
		expect(readBreadcrumbs()).toStrictEqual('...>...')
	})

	it('renders mixed sync and async breadcrumbs', async () => {
		render(Breadcrumbs, {
			props: { entries: ['A', Promise.resolve([Promise.resolve(['B', 'C']), 'D']), 'E'] }
		})

		for (const item of ['A', 'B', 'C', 'D', 'E']) {
			expect(await screen.findByText(item)).toBeInTheDocument()
		}
		expect(readBreadcrumbs()).toStrictEqual('A>B>C>D>E')
	})

	it('renders separators between items except last', async () => {
		const entries = ['A', 'B', 'C']
		render(Breadcrumbs, { props: { entries } })
		const separators = screen.getAllByRole('separator', { hidden: true })
		expect(separators).toHaveLength(entries.length - 1)
	})
})

function readBreadcrumbs(): string {
	const labels: string[] = []
	const root = screen.getByRole('navigation', { name: 'breadcrumb' })
	const items = root.querySelectorAll('li')
	for (const item of items) {
		if (item.dataset.testid === 'breadcrumb-ellipsis') {
			labels.push('...')
		} else if (item.role === 'separator') {
			labels.push('>')
		} else {
			labels.push(item.textContent ?? 'null')
		}
	}
	return labels.join('')
}
