import { expect, Page, test } from '@playwright/test'

test('airport list has expected title and initial flights', async ({ page }) => {
	await page.goto('/admin/airports')
	await expect(page.locator('h1')).toContainText('Airports')
	await expect(page.getByText('SIN')).toBeVisible()
})

test('add and delete airport', async ({ page }) => {
	await addAirport(page, 'AAA')
	await gotoAdminAirportPage(page, 'AAA')

	// Delete airport
	const deleteButton = page.getByRole('button', { name: 'Delete airport' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page).toHaveURL(/\/admin\/airports$/)
	await expect(page.getByText('AAA')).not.toBeVisible()
})

test('airport detail page', async ({ page }) => {
	await page.goto('/admin/airports')
	await gotoAdminAirportPage(page, 'SFO')

	await expect(page.getByRole('heading', 'SFO')).toBeVisible()
	await expect(page.getByTestId('flights-to-from-airport')).toContainText('UA1 SFO–SIN')
})

async function addAirport(page: Page, iataCode: string): Promise<void> {
	await page.goto('/admin/airports')

	// Add airport
	const input = page.getByPlaceholder('IATA code')
	await input.fill(iataCode)
	const addButton = page.getByRole('button', { name: 'Add' })
	await addButton.click()
	await expect(page.getByText(iataCode)).toBeVisible()
}

async function gotoAdminAirportPage(page: Page, iataCode: string): Promise<void> {
	await page.goto('/admin/airports')
	await page.getByRole('link', { name: iataCode }).click()
}
