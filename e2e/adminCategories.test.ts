import { expect, test } from '@playwright/test'

test('category list has expected title and initial products', async ({ page }) => {
	await page.goto('/admin/categories')
	await expect(page.locator('h1')).toContainText('Categories')
	await expect(page.getByText('Cookware')).toBeVisible()
})

test('add and delete category', async ({ page }) => {
	await page.goto('/admin/categories')

	// Add category
	const input = page.getByPlaceholder('New category...')
	await input.fill('Test Category')
	const addButton = page.getByRole('button', { name: 'Add' })
	await addButton.click()
	await expect(page.getByText('Test Category')).toBeVisible()

	// Delete category
	await page.getByRole('link', { name: 'Test Category' }).click()
	const deleteButton = page.getByRole('button', { name: 'Delete' })
	page.on('dialog', (dialog) => dialog.accept())
	await deleteButton.click()
	await expect(page.getByText('Test Category')).not.toBeVisible()
})
