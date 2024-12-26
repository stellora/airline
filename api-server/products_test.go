package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/shop/api-server/api"
)

func TestDeleteAllProducts(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Title: "Product 1"},
		{Title: "Product 2"},
	}

	resp, err := handler.DeleteAllProducts(ctx, api.DeleteAllProductsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllProducts204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkProductTitles(t, handler, []string{})
}

func TestListProducts(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Title: "Product 1"},
		{Title: "Product 2"},
	}

	resp, err := handler.ListProducts(ctx, api.ListProductsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListProducts200JSONResponse{
		api.Product{Title: "Product 1"},
		api.Product{Title: "Product 2"},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestListProductsByCategory(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Id: "1", Title: "Product 1"},
		{Id: "2", Title: "Product 2"},
		{Id: "3", Title: "Product 3"},
	}
	categories = []api.Category{
		{Id: "A", Title: "Category A"},
	}
	productCategoryMemberships = []productCategoryMembership{
		{product: "1", category: "A"},
		{product: "2", category: "A"},
	}

	resp, err := handler.ListProductsByCategory(ctx, api.ListProductsByCategoryRequestObject{
		CategoryId: "A",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListProductsByCategory200JSONResponse{
		ProductsInCategory:    []api.Product{{Id: "1", Title: "Product 1"}, {Id: "2", Title: "Product 2"}},
		ProductsNotInCategory: []api.Product{{Id: "3", Title: "Product 3"}},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateProduct(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{}

	resp, err := handler.CreateProduct(ctx, api.CreateProductRequestObject{
		Body: &api.CreateProductJSONRequestBody{
			Title: "New Product",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateProduct201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkProductTitles(t, handler, []string{"New Product"})
}

func TestDeleteProduct(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Id: "1", Title: "Product 1"},
		{Id: "2", Title: "Product 2"},
	}

	resp, err := handler.DeleteProduct(ctx, api.DeleteProductRequestObject{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteProduct204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkProductTitles(t, handler, []string{"Product 2"})
}

func checkProductTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListProducts(context.Background(), api.ListProductsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	products := resp.(api.ListProducts200JSONResponse)
	if len(products) != len(want) {
		t.Errorf("got %d products, want %d", len(products), len(want))
	}
	for i, category := range products {
		if category.Title != want[i] {
			t.Errorf("got title %q, want %q", category.Title, want[i])
		}
	}
}
