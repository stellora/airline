package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/stellora/airline/api-server/api"
)

var (
	products = []api.Product{
		{Title: "Fork"},
		{Title: "Spoon"},
		{Title: "Knife"},
		{Title: "Cast-Iron Pan"},
		{Title: "Baking Sheet"},
		{Title: "Cutting Board"},
		{Title: "Tomato"},
		{Title: "Zucchini"},
		{Title: "Avocado"},
	}
)

func getProduct(id string) *api.Product {
	for i := range products {
		if products[i].Id == id {
			return &products[i]
		}
	}
	return nil
}

func init() {
	for i := range products {
		products[i].Id = strconv.Itoa(i + 1)
	}
}

func (h *Handler) DeleteAllProducts(ctx context.Context, request api.DeleteAllProductsRequestObject) (api.DeleteAllProductsResponseObject, error) {
	products = []api.Product{}
	productCategoryMemberships = nil
	return api.DeleteAllProducts204Response{}, nil
}

func (h *Handler) GetProduct(ctx context.Context, request api.GetProductRequestObject) (api.GetProductResponseObject, error) {
	product := getProduct(request.Id)
	if product == nil {
		return &api.GetProduct404Response{}, nil
	}
	populateProductCategories(product)
	return api.GetProduct200JSONResponse(*product), nil
}

func populateProductCategories(product *api.Product) {
	categories := []api.Category{}
	for _, membership := range productCategoryMemberships {
		if membership.product == product.Id {
			if category := getCategory(membership.category); category != nil {
				categories = append(categories, *category)
			}
		}
	}
	product.Categories = &categories
}

func (h *Handler) ListProducts(ctx context.Context, request api.ListProductsRequestObject) (api.ListProductsResponseObject, error) {
	productsWithCategories := products
	for i := range productsWithCategories {
		populateProductCategories(&productsWithCategories[i])
	}
	return api.ListProducts200JSONResponse(productsWithCategories), nil
}

func (h *Handler) ListProductsByCategory(ctx context.Context, request api.ListProductsByCategoryRequestObject) (api.ListProductsByCategoryResponseObject, error) {
	category := request.CategoryId

	productsInCategory := []api.Product{}
	for _, product := range products {
		for _, membership := range productCategoryMemberships {
			if membership.product == product.Id && membership.category == category {
				productsInCategory = append(productsInCategory, product)
				break
			}
		}
	}

	productsNotInCategory := []api.Product{}
	for _, product := range products {
		inCategory := false
		for _, membership := range productCategoryMemberships {
			if membership.product == product.Id && membership.category == category {
				inCategory = true
				break
			}
		}
		if !inCategory {
			productsNotInCategory = append(productsNotInCategory, product)
		}
	}

	return api.ListProductsByCategory200JSONResponse{
		ProductsInCategory:    productsInCategory,
		ProductsNotInCategory: productsNotInCategory,
	}, nil
}

func (h *Handler) CreateProduct(ctx context.Context, request api.CreateProductRequestObject) (api.CreateProductResponseObject, error) {
	if request.Body.Title == "" {
		return nil, fmt.Errorf("title must not be empty")
	}

	for _, product := range products {
		if product.Title == request.Body.Title {
			return nil, fmt.Errorf("title must be unique across all products")
		}
	}

	newProduct := api.Product{
		Id:      strconv.Itoa(len(products) + 1),
		Title:   request.Body.Title,
		Starred: false,
	}
	products = append(products, newProduct)

	return api.CreateProduct201Response{}, nil
}

func (h *Handler) DeleteProduct(ctx context.Context, request api.DeleteProductRequestObject) (api.DeleteProductResponseObject, error) {
	// Find and remove the product
	for i, product := range products {
		if product.Id == request.Id {
			products = append(products[:i], products[i+1:]...)
			break
		}
	}

	// Remove all category memberships of this product
	newMemberships := []productCategoryMembership{}
	for _, membership := range productCategoryMemberships {
		if membership.product != request.Id {
			newMemberships = append(newMemberships, membership)
		}
	}

	productCategoryMemberships = newMemberships

	return api.DeleteProduct204Response{}, nil
}

func (h *Handler) SetProductStarred(ctx context.Context, request api.SetProductStarredRequestObject) (api.SetProductStarredResponseObject, error) {
	for i := range products {
		if products[i].Id == request.Id {
			products[i].Starred = request.Body.Starred
			return api.SetProductStarred204Response{}, nil
		}
	}
	return nil, fmt.Errorf("product with id %q not found", request.Id)
}
