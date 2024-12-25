import { expect, test } from '@playwright/test'

test('product list has expected title and initial products', async ({ page }) => {
	await page.goto('/products')
	await expect(page.locator('h1')).toContainText('Products')
	await expect(page.getByText('Avocado')).toBeVisible()
})

test('add and delete product', async ({ page }) => {
	await page.goto('/products')

	// Add product
	const input = page.getByPlaceholder('New product...')
	await input.fill('Test Product')
	const addButton = page.getByRole('button', { name: 'Add' })
	await addButton.click()
	await expect(page.getByText('Test Product')).toBeVisible()

	// Delete product
	const productItem = page.locator('li', { hasText: 'Test Product' })
	const deleteButton = productItem.getByRole('button', { name: 'Delete' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page.getByText('Test Product')).not.toBeVisible()
})

test('star product', async ({ page }) => {
	await page.goto('/products')

	const avocadoProduct = page.locator('li', { hasText: 'Avocado' })

	// Star product
	const starButton = avocadoProduct.getByRole('button', { name: /Star|Unstar/ })
	await starButton.click()
	await expect(starButton).toHaveText('Unstar')
	await expect(avocadoProduct).toContainText('⭐')

	// Unstar product
	await starButton.click()
	await expect(starButton).toHaveText('Star')
	await expect(avocadoProduct).not.toContainText('⭐')
})
