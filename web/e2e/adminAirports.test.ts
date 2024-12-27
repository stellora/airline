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
	await page.getByRole('link', { name: 'AMS' }).click()

	await expect(page.getByText('AMS')).toBeVisible()

	// Add flight to airport
	const select = page.locator('select[name="flight"]')
	await select.selectOption({ label: 'Fork' })
	await page.getByRole('button', { name: 'Add flight to airport' }).click()
	await expect(page.getByTestId('flights-in-airport')).toContainText('Fork')

	// Remove flight from airport
	const removeButton = page
		.getByTestId('flights-in-airport')
		.locator('li', { hasText: 'Fork' })
		.getByRole('button', { name: 'Remove from airport' })
	page.on('dialog', (dialog) => dialog.accept())
	await removeButton.click()
	await expect(page.getByTestId('flights-by-airport')).not.toContainText('Fork')
})
