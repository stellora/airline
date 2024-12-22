<script lang="ts">
		import { todos } from '$lib/stores/todoStore';
		import type { Product } from '$lib/types';

		let { todo }: { todo: Product } = $props();
</script>

<div class="flex items-center gap-4">
		<input
				type="checkbox"
				checked={todo.done}
				onchange={() => todos.toggle(todo.id)}
		/>
		<span class={todo.done ? 'line-through text-gray-500' : ''}>{todo.description}</span>
		<button 
			onclick={() => todos.remove(todo.id)}
			class="ml-auto bg-red-500 text-white rounded cursor-pointer"
		>Delete</button>
</div>