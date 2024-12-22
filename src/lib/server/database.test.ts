import { beforeEach, describe, expect, it } from 'vitest';
import { createProduct, deleteAllProducts, deleteProduct, listProducts } from './database';

describe('database', () => {
	beforeEach(() => {
		deleteAllProducts();
	});

	describe('listProducts', () => {
		it('returns empty array when no products exist', () => {
			expect(listProducts()).toEqual([]);
		});

		it('returns array of products when products exist', () => {
			createProduct('Test Product');
			const products = listProducts();
			expect(products).toHaveLength(1);
			expect(products[0].title).toBe('Test Product');
			expect(products[0].id).toBeDefined();
		});
	});

	describe('createProduct', () => {
		it('creates a product with given title', () => {
			createProduct('Test Product');
			const products = listProducts();
			expect(products[0].title).toBe('Test Product');
		});

		it('generates unique id for each product', () => {
			createProduct('Product 1');
			createProduct('Product 2');
			const products = listProducts();
			expect(products[0].id).not.toBe(products[1].id);
		});

		it('throws error when title is empty', () => {
			expect(() => createProduct('')).toThrow('title must not be empty');
		});

		it('throws error when title is duplicate', () => {
			createProduct('Test Product');
			expect(() => createProduct('Test Product')).toThrow('title must be unique across all products');
		});
	});

	describe('deleteProduct', () => {
		it('removes product with given id', () => {
			createProduct('Product 1');
			createProduct('Product 2');
			const products = listProducts();
			deleteProduct(products[0].id);
			expect(listProducts()).toHaveLength(1);
		});

		it('does nothing when id does not exist', () => {
			createProduct('Test Product');
			deleteProduct('non-existent-id');
			expect(listProducts()).toHaveLength(1);
		});
	});

	describe('deleteAllProducts', () => {
		it('removes all products', () => {
			createProduct('Product 1');
			createProduct('Product 2');
			deleteAllProducts();
			expect(listProducts()).toHaveLength(0);
		});
	});
});
