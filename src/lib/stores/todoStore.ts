import type { Product } from '$lib/types';
import { writable } from 'svelte/store';

function createTodoStore() {
    const { subscribe, set, update } = writable<Product[]>([]);

    return {
        subscribe,
        add: (text: string) => update(todos => [
            ...todos,
            { id: crypto.randomUUID(), description: text, done: false }
        ]),
        remove: (id: string) => update(todos => 
            todos.filter(todo => todo.id !== id)
        ),
        toggle: (id: string) => update(todos => 
            todos.map(todo => 
                todo.id === id 
                    ? { ...todo, done: !todo.done }
                    : todo
            )
        ),
        clear: () => set([])
    };
}

export const todos = createTodoStore();