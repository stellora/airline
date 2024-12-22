import type { Todo } from '$lib/types';
import { writable } from 'svelte/store';

function createTodoStore() {
    const { subscribe, set, update } = writable<Todo[]>([]);

    return {
        subscribe,
        add: (text: string) => update(todos => [
            ...todos,
            { id: crypto.randomUUID(), text, completed: false }
        ]),
        remove: (id: string) => update(todos => 
            todos.filter(todo => todo.id !== id)
        ),
        toggle: (id: string) => update(todos => 
            todos.map(todo => 
                todo.id === id 
                    ? { ...todo, completed: !todo.completed }
                    : todo
            )
        ),
        clear: () => set([])
    };
}

export const todos = createTodoStore();