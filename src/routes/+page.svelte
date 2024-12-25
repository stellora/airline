<script lang="ts">
	import ProductForm from './ProductForm.svelte'
	import ProductList from './ProductList.svelte'

	let { data, form } = $props()
</script>

<div class="flex flex-col gap-10 items-stretch w-full max-w-[800px] mx-auto">
	<h1 class="text-4xl font-bold">Products</h1>
	<ProductForm {form} />
	<ProductList products={data.products} />
</div>
