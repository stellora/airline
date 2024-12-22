import type { Product } from '$lib/types';

const todos: Product[] = [];

export function getProducts() {
	return todos;
}

export function createProduct(description: string):void {
	if (description === '') {
		throw new Error('todo must have a description');
	}

	if (todos.find((todo) => todo.description === description)) {
		throw new Error('todos must be unique');
	}

	todos.push({
		id: crypto.randomUUID(),
		description,
		done: false
	});
}

export function deleteProduct(id: string):void {
	const index = todos.findIndex((todo) => todo.id === id);
	if (index !== -1) {
		todos.splice(index, 1);
	}
}