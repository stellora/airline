<script lang="ts">
	import ProductForm from './product-form.svelte'
	import ProductList from './product-list.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-4 items-stretch w-full">
	<h1 class="text-2xl font-bold">Products</h1>
	<ProductForm {form} />
	<ProductList products={data.products} />
</div>
