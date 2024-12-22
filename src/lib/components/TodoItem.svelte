<script lang="ts">
		import { todos } from '$lib/stores/todoStore';
		import type { Todo } from '$lib/types';

		let { todo }: { todo: Todo } = $props();
</script>

<div class="flex items-center gap-4">
		<input
				type="checkbox"
				checked={todo.completed}
				onchange={() => todos.toggle(todo.id)}
		/>
		<span class={todo.completed ? 'line-through text-gray-500' : ''}>{todo.text}</span>
		<button 
			onclick={() => todos.remove(todo.id)}
			class="ml-auto bg-red-500 text-white rounded cursor-pointer"
		>Delete</button>
</div>