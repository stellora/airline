import { expect, test } from '@playwright/test'

test('airport list has expected title and initial flights', async ({ page }) => {
	await page.goto('/admin/airports')
	await expect(page.locator('h1')).toContainText('Airports')
	await expect(page.getByText('AMS')).toBeVisible()
})

test('add and delete airport', async ({ page }) => {
	await page.goto('/admin/airports')

	// Add airport
	const input = page.getByPlaceholder('IATA code')
	await input.fill('AAA')
	const addButton = page.getByRole('button', { name: 'Add' })
	await addButton.click()
	await expect(page.getByText('AAA')).toBeVisible()

	// Delete airport
	await page.getByRole('link', { name: 'AAA' }).click()
	const deleteButton = page.getByRole('button', { name: 'Delete' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page.getByText('AAA')).not.toBeVisible()
})

test('airport detail page', async ({ page }) => {
	await page.goto('/admin/airports')
	await page.getByRole('link', { name: 'SFO' }).click()

	await expect(page.getByText('SFO')).toBeVisible()
	await expect(page.getByTestId('flights-to-from-airport')).toContainText('UA1')
})
