package main

import (
	"context"
	"strconv"

	"github.com/stellora/shop/api-server/api"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct{}

var (
	categories = []api.Category{
		{Title: "Silverware"},
		{Title: "Cookware"},
		{Title: "Vegetables"},
	}
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
	productCategoryMemberships = []productCategoryMembership{
		{product: "1", category: "1"},
		{product: "2", category: "1"},
		{product: "3", category: "1"},
		{product: "4", category: "2"},
		{product: "5", category: "2"},
		{product: "6", category: "2"},
		{product: "7", category: "3"},
		{product: "8", category: "3"},
		{product: "9", category: "3"},
	}
)

type productCategoryMembership struct {
	product, category string
}

func init() {
	for i := range categories {
		categories[i].Id = strconv.Itoa(i + 1)
	}
	for i := range products {
		products[i].Id = strconv.Itoa(i + 1)
	}
}

func (h *Handler) ListCategories(ctx context.Context, request api.ListCategoriesRequestObject) (api.ListCategoriesResponseObject, error) {
	return api.ListCategories200JSONResponse(categories), nil
}

func (h *Handler) CreateCategory(ctx context.Context, request api.CreateCategoryRequestObject) (api.CreateCategoryResponseObject, error) {
	return api.CreateCategory201Response{}, nil
}

func (h *Handler) ListProductsByCategory(ctx context.Context, request api.ListProductsByCategoryRequestObject) (api.ListProductsByCategoryResponseObject, error) {
	return api.ListProductsByCategory200JSONResponse{}, nil
}

func (h *Handler) DeleteCategory(ctx context.Context, request api.DeleteCategoryRequestObject) (api.DeleteCategoryResponseObject, error) {
	return api.DeleteCategory204Response{}, nil
}

func (h *Handler) HealthCheck(ctx context.Context, request api.HealthCheckRequestObject) (api.HealthCheckResponseObject, error) {
	ok := true
	return api.HealthCheck200JSONResponse{Ok: &ok}, nil
}

func (h *Handler) DeleteAllProducts(ctx context.Context, request api.DeleteAllProductsRequestObject) (api.DeleteAllProductsResponseObject, error) {
	return api.DeleteAllProducts204Response{}, nil
}

func (h *Handler) ListProducts(ctx context.Context, request api.ListProductsRequestObject) (api.ListProductsResponseObject, error) {
	return api.ListProducts200JSONResponse(products), nil
}

func (h *Handler) CreateProduct(ctx context.Context, request api.CreateProductRequestObject) (api.CreateProductResponseObject, error) {
	return api.CreateProduct201Response{}, nil
}

func (h *Handler) DeleteProduct(ctx context.Context, request api.DeleteProductRequestObject) (api.DeleteProductResponseObject, error) {
	return api.DeleteProduct204Response{}, nil
}

func (h *Handler) SetProductStarred(ctx context.Context, request api.SetProductStarredRequestObject) (api.SetProductStarredResponseObject, error) {
	return api.SetProductStarred200Response{}, nil
}

func (h *Handler) UpdateProductCategoryMembership(ctx context.Context, request api.UpdateProductCategoryMembershipRequestObject) (api.UpdateProductCategoryMembershipResponseObject, error) {
	return api.UpdateProductCategoryMembership200Response{}, nil
}

var _ api.StrictServerInterface = (*Handler)(nil)
