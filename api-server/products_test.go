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

func TestGetProduct(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Id: "1", Title: "Product 1"},
		{Id: "2", Title: "Product 2"},
	}
	categories = []api.Category{
		{Id: "A", Title: "Category A"},
	}
	productCategoryMemberships = []productCategoryMembership{
		{product: "1", category: "A"},
	}

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetProduct(ctx, api.GetProductRequestObject{
			Id: "1",
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetProduct200JSONResponse{
			Id:         "1",
			Title:      "Product 1",
			Categories: &[]api.Category{{Id: "A", Title: "Category A"}},
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetProduct(ctx, api.GetProductRequestObject{
			Id: "999",
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetProduct404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})
}

func TestListProducts(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Id: "1", Title: "Product 1"},
		{Id: "2", Title: "Product 2"},
	}
	categories = []api.Category{
		{Id: "A", Title: "Category A"},
	}
	productCategoryMemberships = []productCategoryMembership{
		{product: "1", category: "A"},
	}

	resp, err := handler.ListProducts(ctx, api.ListProductsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListProducts200JSONResponse{
		api.Product{
			Id:         "1",
			Title:      "Product 1",
			Categories: &[]api.Category{{Id: "A", Title: "Category A"}},
		},
		api.Product{Id: "2", Title: "Product 2", Categories: &[]api.Category{}},
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

func TestSetProductStarred(t *testing.T) {
	ctx, handler := handlerTest(t)
	products = []api.Product{
		{Id: "1", Title: "Product 1", Starred: false},
		{Id: "2", Title: "Product 2", Starred: false},
	}

	resp, err := handler.SetProductStarred(ctx, api.SetProductStarredRequestObject{
		Id: "1",
		Body: &api.SetProductStarredJSONRequestBody{
			Starred: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.SetProductStarred204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	// Verify the product was actually starred
	listResp, err := handler.ListProducts(ctx, api.ListProductsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	products := listResp.(api.ListProducts200JSONResponse)
	for _, p := range products {
		if p.Id == "1" && !p.Starred {
			t.Error("Product 1 should be starred")
		}
		if p.Id == "2" && p.Starred {
			t.Error("Product 2 should not be starred")
		}
	}
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
