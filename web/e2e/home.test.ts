import { expect, test } from '@playwright/test'

test('homepage has initial flights', async ({ page }) => {
	await page.goto('/')
	await expect(page.locator('h1')).toContainText('Book flights')
	await expect(page.getByText('UA1')).toBeVisible()
	await expect(page.locator('li', { hasText: 'UA1' })).toContainText('SFO')
})
