package main

import (
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestUpdateProductCategoryMembership(t *testing.T) {
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

	{
		// Add to the category
		resp, err := handler.UpdateProductCategoryMembership(ctx, api.UpdateProductCategoryMembershipRequestObject{
			ProductId:  "2",
			CategoryId: "A",
			Body: &api.UpdateProductCategoryMembershipJSONRequestBody{
				Value: true,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (api.UpdateProductCategoryMembership204Response{}); !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}

		// Verify the membership was actually updated
		listResp, err := handler.ListProductsByCategory(ctx, api.ListProductsByCategoryRequestObject{
			CategoryId: "A",
		})
		if err != nil {
			t.Fatal(err)
		}
		result := listResp.(api.ListProductsByCategory200JSONResponse)
		if want := 2; len(result.ProductsInCategory) != want {
			t.Errorf("got %d products in category, want %d", len(result.ProductsInCategory), want)
		}
	}

	{
		// Remove from the category
		resp, err := handler.UpdateProductCategoryMembership(ctx, api.UpdateProductCategoryMembershipRequestObject{
			ProductId:  "1",
			CategoryId: "A",
			Body: &api.UpdateProductCategoryMembershipJSONRequestBody{
				Value: false,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (api.UpdateProductCategoryMembership204Response{}); !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}

		// Verify the membership was actually updated
		listResp, err := handler.ListProductsByCategory(ctx, api.ListProductsByCategoryRequestObject{
			CategoryId: "A",
		})
		if err != nil {
			t.Fatal(err)
		}
		result := listResp.(api.ListProductsByCategory200JSONResponse)
		if want := 1; len(result.ProductsInCategory) != want {
			t.Errorf("got %d products in category, want %d", len(result.ProductsInCategory), want)
		}
	}
}
