import { expect, test } from '@playwright/test'

test('product list has expected title and initial products', async ({ page }) => {
	await page.goto('/')
	await expect(page.locator('h1')).toContainText('Shop')
	await expect(page.getByText('Avocado')).toBeVisible()
})
