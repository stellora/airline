import { expect, test } from '@playwright/test'

test('flight list has expected title and initial flights', async ({ page }) => {
	await page.goto('/admin/flights')
	await expect(page.locator('h1')).toContainText('Flights')
	await expect(page.getByText('Avocado')).toBeVisible()
})

test('add and delete flight', async ({ page }) => {
	await page.goto('/admin/flights')

	// Add flight
	const input = page.getByPlaceholder('New flight...')
	await input.fill('Test Flight')
	const addButton = page.getByRole('button', { name: 'Add' })
	await addButton.click()
	await expect(page.getByText('Test Flight')).toBeVisible()

	// Delete flight
	const flightItem = page.locator('li', { hasText: 'Test Flight' })
	const deleteButton = flightItem.getByRole('button', { name: 'Delete' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page.getByText('Test Flight')).not.toBeVisible()
})

test('star flight', async ({ page }) => {
	await page.goto('/admin/flights')

	const avocadoFlight = page.locator('li', { hasText: 'Avocado' })

	// Star flight
	const starButton = avocadoFlight.getByRole('button', { name: /Star|Unstar/ })
	await starButton.click()
	await expect(starButton).toHaveText('Unstar')
	await expect(avocadoFlight.getByLabel('Starred')).toBeVisible()

	// Unstar flight
	await starButton.click()
	await expect(starButton).toHaveText('Star')
	await expect(avocadoFlight.getByLabel('Starred')).not.toBeVisible()
})
