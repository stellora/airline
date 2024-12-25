<script lang="ts">
	import { page } from '$app/state'

	let { children } = $props()
</script>

<nav class="bg-gray-800 p-4">
	<div class="container mx-auto">
		<ul class="flex space-x-4">
			<li>
				<a
					href="/admin/products"
					class="text-white hover:text-gray-300 {page.url.pathname === '/admin/products'
						? 'font-bold'
						: ''}"
				>
					Products
				</a>
			</li>
		</ul>
	</div>
</nav>

<main class="container mx-auto p-4">
	{@render children()}
</main>
