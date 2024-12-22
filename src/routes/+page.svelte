<script lang="ts">
	import TodoItem from '$lib/components/TodoItem.svelte';
	import { todos } from '$lib/stores/todoStore';
	
	let newTodoText = $state('');

	function handleSubmit() {
			if (newTodoText.trim()) {
					todos.add(newTodoText);
					newTodoText = '';
			}
	}
</script>

<svelte:head>
	<title>TODOs</title>
</svelte:head>

<main class="max-w-[600px] mx-auto my-8 px-4 flex flex-col gap-6">
	<h1 class="text-4xl font-bold text-center">TODOs</h1>
	
	<form onsubmit={handleSubmit} class="flex gap-4">
			<input
					type="text"
					bind:value={newTodoText}
					placeholder="Add a new todo..."
					class="flex-1"
			/>
			<button type="submit" class="bg-green-500 text-white cursor-pointer">Add</button>
	</form>

	<div class="flex flex-col gap-2">
			{#each $todos as todo (todo.id)}
					<TodoItem {todo} />
			{/each}
	</div>

	{#if $todos.length > 0}
			<button class="cursor-pointer" onclick={() => todos.clear()}>Clear All</button>
	{/if}
</main>
