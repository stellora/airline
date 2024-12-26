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

test('category detail page', async ({ page }) => {
	await page.goto('/admin/categories')
	await page.getByRole('link', { name: 'Cookware' }).click()

	await expect(page.getByText('Cookware')).toBeVisible()

	// Add product to category
	const select = page.locator('select[name="product"]')
	await select.selectOption({ label: 'Fork' })
	await page.getByRole('button', { name: 'Add product to category' }).click()
	await expect(page.getByTestId('products-in-category')).toContainText('Fork')

	// Remove product from category
	const removeButton = page
		.getByTestId('products-in-category')
		.locator('li', { hasText: 'Fork' })
		.getByRole('button', { name: 'Remove from category' })
	page.on('dialog', (dialog) => dialog.accept())
	await removeButton.click()
	await expect(page.getByTestId('products-in-category')).not.toContainText('Fork')
})
