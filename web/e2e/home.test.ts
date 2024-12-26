import { expect, test } from '@playwright/test'

test('homepage has initial products', async ({ page }) => {
	await page.goto('/')
	await expect(page.locator('h1')).toContainText('Airline')
	await expect(page.getByText('Avocado')).toBeVisible()
	await expect(page.locator('li', { hasText: 'Avocado' })).toContainText('Vegetables')
})
