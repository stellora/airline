import { expect, Page, test } from '@playwright/test'

test('flight list has expected title and initial flights', async ({ page }) => {
	await page.goto('/admin/flight-schedules')
	await expect(page.locator('h1')).toContainText('Flights')
	await expect(page.getByText('UA1 SFO–SIN')).toBeVisible()
})

test('add and delete flight', async ({ page }) => {
	await addFlight(page, 'ST123', 'SFO', 'SIN')
	await gotoAdminFlightPage(page, 'ST123')

	// Delete flight
	const deleteButton = page.getByRole('button', { name: 'Delete flight' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page).toHaveURL(/\/admin\/flight-schedules$/)
	await expect(page.getByText('ST123 SFO–LAX')).not.toBeVisible()
})

test('publish flight', async ({ page }) => {
	await addFlight(page, 'ST456', 'SIN', 'SFO')
	await gotoAdminFlightPage(page, 'ST456')

	// Initially unpublished
	const flightNumber = page.locator('span', { hasText: 'ST456' })
	await expect(flightNumber).toHaveCSS('text-decoration-style', 'dotted')

	// Publish flight
	const publishButton = page.getByRole('button', { name: /Publish|Unpublish/ })
	await expect(publishButton).toHaveText('Publish')
	await publishButton.click()
	await expect(flightNumber).not.toHaveCSS('text-decoration-style', 'dotted')

	// Unpublish flight
	await expect(publishButton).toHaveText('Unpublish')
	await publishButton.click()
	await expect(flightNumber).toHaveCSS('text-decoration-style', 'dotted')
})

async function addFlight(
	page: Page,
	number: string,
	originIataCode: string,
	destinationIataCode: string,
): Promise<void> {
	await page.goto('/admin/flight-schedules')

	// Add flight
	await page.getByLabel('Flight number').fill(number)
	await page.getByLabel('Origin airport IATA code').fill(originIataCode)
	await page.getByLabel('Destination airport IATA code').fill(destinationIataCode)
	const addButton = page.getByRole('button', { name: 'Add flight' })
	await addButton.click()

	await expect(page.getByText(`${number} ${originIataCode}–${destinationIataCode}`)).toBeVisible()
}

async function gotoAdminFlightPage(page: Page, number: string): Promise<void> {
	await page.goto('/admin/flight-schedules')
	await page.getByRole('link', { name: number }).click()
}
